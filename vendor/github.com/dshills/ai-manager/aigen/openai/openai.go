package openai

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

const AIName = "openai"

const (
	roleSystem    = "system"
	roleAssistant = "assistant"
	roleUser      = "user"
)

const (
	keyTemperature = "temperature"
	keyMaxTokens   = "max_tokens"
)

type Generator struct {
	model   string
	apiKey  string
	baseURL string
	tools   map[string]aitool.Tool
}

func New(model, apiKey, baseURL string) ai.Generator {
	return &Generator{model: model, apiKey: apiKey, baseURL: baseURL, tools: make(map[string]aitool.Tool)}
}

func (g *Generator) Model() string {
	return g.model
}

func (g *Generator) getTemp(meta []ai.Meta) int {
	for _, m := range meta {
		if m.Key == keyTemperature {
			val, err := strconv.Atoi(m.Value)
			if err != nil || val < 0 || val > 2 {
				return -1
			}
			return val
		}
	}
	return -1
}

func (g *Generator) getMaxTokens(meta []ai.Meta) int {
	for _, m := range meta {
		if m.Key == keyMaxTokens {
			val, err := strconv.Atoi(m.Value)
			if err != nil {
				return -1
			}
			return val
		}
	}
	return -1
}

func (g *Generator) NewRequest(conversation ai.Conversation, meta []ai.Meta, tools []aitool.Tool) CreateRequest {
	frags := []MessageCallFrag{}
	for _, m := range conversation {
		frags = append(frags, MessageCallFrag{Role: m.Role, Content: m.Text})
	}
	chatReq := CreateRequest{
		Model:    g.model,
		Messages: frags,
	}
	maxToks := g.getMaxTokens(meta)
	if maxToks > 0 {
		chatReq.MaxTokens = &maxToks
	}
	temp := g.getTemp(meta)
	if temp > -1 { // 0 - 2 float
		chatReq.Temperature = temp
	}
	if len(tools) > 0 {
		chatReq.Tools = tools
		chatReq.ToolChoice = "auto"

		// Save tools for future calls
		for _, t := range tools {
			g.tools[t.Name()] = t
		}
	}

	return chatReq
}

func (g *Generator) Generate(conversation ai.Conversation, meta []ai.Meta, tools []aitool.Tool) (*ai.GeneratorResponse, error) {
	chatReq := g.NewRequest(conversation, meta, tools)
	byts, err := json.MarshalIndent(&chatReq, "", "\t")
	if err != nil {
		return nil, fmt.Errorf("openai.Generator: %w", err)
	}

	response := ai.GeneratorResponse{}

	start := time.Now()
	resp, err := completion(g.apiKey, g.baseURL, bytes.NewReader(byts))
	if err != nil {
		return nil, fmt.Errorf("openai.Generator: %w", err)
	}

	response.Elapsed = time.Since(start)
	response.Usage.PromptTokens = resp.Usage.PromptTokens
	response.Usage.CompletionTokens = resp.Usage.CompletionTokens
	response.Usage.TotalTokens = resp.Usage.TotalTokens
	response.Message.Role = roleAssistant
	response.Message.Text = resp.Choices[0].Message.Content
	response.FinishReason = resp.Choices[0].FinishReason
	for _, call := range resp.Choices[0].Message.ToolCalls {
		tc := ai.ToolCall{
			ID:   call.ID,
			Type: call.Type,
			Name: call.Function.Name,
			Args: call.Function.Arguments,
		}
		response.ToolCalls = append(response.ToolCalls, tc)
	}

	return &response, nil
}

func completion(apiKey, baseURL string, reader io.Reader) (*ChatResp, error) {
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
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("thread.completion: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		bdy, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("thread.completion: %v %v %v", resp.StatusCode, resp.Status, string(bdy))
	}

	chatResp := ChatResp{}
	err = json.NewDecoder(resp.Body).Decode(&chatResp)
	if err != nil {
		return nil, fmt.Errorf("thread.completion: %w", err)
	}
	if len(chatResp.Choices) == 0 {
		return nil, fmt.Errorf("thread.completion: No data returned")
	}

	return &chatResp, nil
}
