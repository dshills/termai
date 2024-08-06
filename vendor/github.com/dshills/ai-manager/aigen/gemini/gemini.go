package gemini

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/dshills/ai-manager/ai"
	"github.com/dshills/ai-manager/aitool"
)

const AIName = "gemini"

const (
	geminiEP = "/models/%%MODEL%%:generateContent?key=%%APIKEY%%"
)

type Generator struct {
	model   string
	baseURL string
	apiKey  string
	tools   map[string]aitool.Tool
}

func New(model, apiKey, baseURL string) ai.Generator {
	return &Generator{model: model, apiKey: apiKey, baseURL: baseURL, tools: make(map[string]aitool.Tool)}
}

func (g *Generator) Model() string {
	return g.model
}

func (g *Generator) Generate(conversation ai.Conversation, _ []ai.Meta, _ []aitool.Tool) (*ai.GeneratorResponse, error) {
	conlist := []Content{}
	for _, m := range conversation {
		con := Content{Role: m.Role, Parts: []Part{{Text: m.Text}}}
		conlist = append(conlist, con)
	}
	req := Request{Contents: conlist}
	body, err := json.Marshal(&req)
	if err != nil {
		return nil, fmt.Errorf("gemini.Generator: %w", err)
	}

	response := ai.GeneratorResponse{}

	start := time.Now()
	resp, err := completion(g.model, g.apiKey, g.baseURL, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("gemini.Generator: %w", err)
	}

	response.Elapsed = time.Since(start)

	for _, can := range resp.Candidates {
		response.Usage.TotalTokens += int64(can.TokenCount)
	}

	response.Message.Role = resp.Candidates[0].Content.Role
	response.Message.Text = resp.Candidates[0].Content.Parts[0].Text

	return &response, nil
}

func completion(model, apiKey, baseURL string, reader io.Reader) (*Response, error) {
	ep := fmt.Sprintf("%v%v", baseURL, geminiEP)
	ep = strings.Replace(ep, "%%MODEL%%", model, 1)
	ep = strings.Replace(ep, "%%APIKEY%%", apiKey, 1)

	client := http.Client{}
	req, err := http.NewRequest(http.MethodPost, ep, reader)
	if err != nil {
		return nil, fmt.Errorf("completion: %w", err)
	}

	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("gemini.completion: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("gemini.completion: %v %v", resp.StatusCode, resp.Status)
	}

	chatResp := Response{}
	err = json.NewDecoder(resp.Body).Decode(&chatResp)
	if err != nil {
		return nil, fmt.Errorf("gemini.completion: %w", err)
	}
	if len(chatResp.Candidates) == 0 {
		return nil, fmt.Errorf("gemini.completion: No data returned")
	}
	if len(chatResp.Candidates[0].Content.Parts) == 0 {
		return nil, fmt.Errorf("gemini.completion: No data returned")
	}

	return &chatResp, nil
}
