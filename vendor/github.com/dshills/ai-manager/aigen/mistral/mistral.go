package mistral

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

const chatEP = "/chat/completions"

const AIName = "mistral"

const (
	keyTemperature = "temperature"
	keyMaxTokens   = "max_tokens"
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

func (g *Generator) getTemperature(meta []ai.Meta) float64 {
	for _, m := range meta {
		if m.Key == keyTemperature {
			temp, err := strconv.ParseFloat(m.Value, 64)
			if err != nil {
				return 0.2
			}
			return temp
		}
	}
	return 0.2
}

func (g *Generator) getMaxTokens(meta []ai.Meta) int {
	for _, m := range meta {
		if m.Key == keyTemperature {
			maxT, err := strconv.Atoi(m.Value)
			if err != nil {
				return -1
			}
			return maxT
		}
	}
	return -1
}

func (g *Generator) Generate(conversation ai.Conversation, meta []ai.Meta, _ []aitool.Tool) (*ai.GeneratorResponse, error) {
	messages := []Message{}
	for _, m := range conversation {
		msg := Message{Role: m.Role, Content: m.Text}
		messages = append(messages, msg)
	}
	req := Request{
		Model:       g.model,
		Messages:    messages,
		Stream:      false,
		SafePrompt:  false,
		Temperature: g.getTemperature(meta),
	}
	maxTok := g.getMaxTokens(meta)
	if maxTok > 0 {
		req.MaxTokens = maxTok
	}
	body, err := json.Marshal(&req)
	if err != nil {
		return nil, fmt.Errorf("mistral.Generator: %w", err)
	}

	response := ai.GeneratorResponse{}

	start := time.Now()
	resp, err := completion(g.apiKey, g.baseURL, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("mistral.Generator: %w", err)
	}

	response.Elapsed = time.Since(start)
	response.Usage.PromptTokens = int64(resp.Usage.PromptTokens)
	response.Usage.CompletionTokens = int64(resp.Usage.CompletionTokens)
	response.Usage.TotalTokens = int64(resp.Usage.TotalTokens)
	response.Message.Role = resp.Choices[0].Message.Role
	response.Message.Text = resp.Choices[0].Message.Content
	response.FinishReason = resp.Choices[0].FinishReason

	return &response, nil
}

func completion(apiKey, baseURL string, reader io.Reader) (*Response, error) {
	ep, err := url.JoinPath(baseURL, chatEP)
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	req, err := http.NewRequest(http.MethodPost, ep, reader)
	if err != nil {
		return nil, fmt.Errorf("completion: %w", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("mistral.completion: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("mistral.completion: %v %v", resp.StatusCode, resp.Status)
	}

	chatResp := Response{}
	err = json.NewDecoder(resp.Body).Decode(&chatResp)
	if err != nil {
		return nil, fmt.Errorf("mistral.completion: %w", err)
	}
	if len(chatResp.Choices) == 0 {
		return nil, fmt.Errorf("mistral.completion: No data returned")
	}
	if len(chatResp.Choices[0].Message.Content) == 0 {
		return nil, fmt.Errorf("mistral.completion: No data returned")
	}

	return &chatResp, nil
}
