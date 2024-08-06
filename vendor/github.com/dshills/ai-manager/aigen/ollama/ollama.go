package ollama

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/dshills/ai-manager/ai"
	"github.com/dshills/ai-manager/aitool"
)

const chatEP = "api/chat"
const keyTemperature = "temperature"

type Generator struct {
	model   string
	baseURL string
	tools   map[string]aitool.Tool
}

func New(model, baseURL string) ai.Generator {
	return &Generator{model: model, baseURL: baseURL, tools: make(map[string]aitool.Tool)}
}

func (g *Generator) Model() string {
	return g.model
}

// getTemp 0 - 1 float 64
func (g *Generator) getTemp(meta []ai.Meta) float64 {
	for _, m := range meta {
		if m.Key == keyTemperature {
			val, err := strconv.ParseFloat(m.Value, 64)
			if err != nil {
				return -1
			}
			return val
		}
	}
	return -1
}

func (g *Generator) Generate(conversation ai.Conversation, meta []ai.Meta, _ []aitool.Tool) (*ai.GeneratorResponse, error) {
	chatReq := ChatRequest{
		Model: g.model,
	}
	temp := g.getTemp(meta)
	if temp > -1 {
		chatReq.Options.Temperature = temp
	}
	chatReq.convConv(conversation)

	byts, err := json.MarshalIndent(&chatReq, "", "\t")
	if err != nil {
		return nil, fmt.Errorf("ollama.Generator: %w", err)
	}

	response := ai.GeneratorResponse{}

	start := time.Now()
	resp, err := completion(g.baseURL, bytes.NewReader(byts))
	if err != nil {
		return nil, fmt.Errorf("ollama.Generator: %w", err)
	}

	response.Elapsed = time.Since(start)
	response.Usage.PromptTokens = int64(resp.PromptEvalCount)
	response.Usage.CompletionTokens = int64(resp.EvalCount)
	response.Usage.TotalTokens = response.Usage.PromptTokens + response.Usage.CompletionTokens
	response.Message.Role = roleAssistant
	response.Message.Text = resp.Message.Content

	return &response, nil
}

func completion(baseURL string, reader io.Reader) (*ChatResponse, error) {
	ep, err := url.JoinPath(baseURL, chatEP)
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	req, err := http.NewRequest(http.MethodPost, ep, reader)
	if err != nil {
		return nil, fmt.Errorf("thread.completion: %w", err)
	}

	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("thread.completion: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("thread.completion: %v %v", resp.StatusCode, resp.Status)
	}

	chatResp := ChatResponse{}
	err = json.NewDecoder(resp.Body).Decode(&chatResp)
	if err != nil {
		return nil, fmt.Errorf("thread.completion: %w", err)
	}
	if len(chatResp.Message.Content) == 0 {
		return nil, fmt.Errorf("thread.completion: No data returned")
	}

	return &chatResp, nil
}
