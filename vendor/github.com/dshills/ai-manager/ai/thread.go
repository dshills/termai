package ai

import (
	"time"

	"github.com/dshills/ai-manager/aitool"
	"github.com/google/uuid"
)

type ThreadData struct {
	ID               string
	Model            string
	CreatedAt        time.Time
	Conversation     Conversation
	CompletionTokens int64
	PromptTokens     int64
	TotalTokens      int64
	Tools            []aitool.Tool
	MetaData         []Meta
}

func NewThreadData(model string) ThreadData {
	return ThreadData{
		ID:    uuid.New().String(),
		Model: model,
	}
}

type Thread interface {
	ID() string
	Conversation() Conversation
	Info() ThreadData
	Converse(query string) (*GeneratorResponse, error)
}

type _thread struct {
	info ThreadData
	gen  Generator
	mgr  *Manager
}

func (t *_thread) ID() string {
	return t.info.ID
}

func (t *_thread) Conversation() Conversation {
	return t.info.Conversation
}

func (t *_thread) Info() ThreadData {
	return t.info
}

func (t *_thread) updateConv(msg Message) {
	t.info.Conversation = append(t.info.Conversation, msg)
}

func (t *_thread) updateUsage(usage Usage) {
	t.info.PromptTokens += usage.PromptTokens
	t.info.CompletionTokens += usage.CompletionTokens
	t.info.TotalTokens += usage.TotalTokens
}

func (t *_thread) Converse(query string) (*GeneratorResponse, error) {
	msg := Message{Role: "user", Text: query}
	t.updateConv(msg)

	resp, err := t.gen.Generate(t.info.Conversation, t.info.MetaData, t.info.Tools)
	if err != nil {
		return nil, err
	}
	t.updateUsage(resp.Usage)
	t.updateConv(resp.Message)
	return resp, nil
}

func newThread(mgr *Manager, info ThreadData, gen Generator) Thread {
	return &_thread{mgr: mgr, info: info, gen: gen}
}
