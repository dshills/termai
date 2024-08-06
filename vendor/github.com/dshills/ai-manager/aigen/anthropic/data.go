package anthropic

import (
	"github.com/dshills/ai-manager/ai"
)

type Request struct {
	Model     string        `json:"model,omitempty"`
	MaxTokens int           `json:"max_tokens,omitempty"`
	Messages  []MessageFrag `json:"messages,omitempty"`
}

func (r *Request) fillMsgs(conversation ai.Conversation) {
	for _, conv := range conversation {
		r.Messages = append(r.Messages, MessageFrag{Role: conv.Role, Content: conv.Text})
	}
}

type MessageFrag struct {
	Role    string `json:"role,omitempty"`
	Content string `json:"content,omitempty"`
}

type Response struct {
	ID      string `json:"id"`
	Content []struct {
		Text  string `json:"text,omitempty"`
		ID    string `json:"id,omitempty"`
		Name  string `json:"name,omitempty"`
		Input struct {
		} `json:"input,omitempty"`
	} `json:"content"`
	Model        string `json:"model"`
	StopReason   string `json:"stop_reason"`
	StopSequence string `json:"stop_sequence"`
	Usage        struct {
		InputTokens  int `json:"input_tokens"`
		OutputTokens int `json:"output_tokens"`
	} `json:"usage"`
}
