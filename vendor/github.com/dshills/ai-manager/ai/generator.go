package ai

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/dshills/ai-manager/aitool"
)

// GeneratorResponse is the response data from a Generator call
type GeneratorResponse struct {
	Elapsed      time.Duration
	Message      Message
	Usage        Usage
	Meta         []Meta
	ToolCalls    []ToolCall
	FinishReason string
}

type ToolCall struct {
	ID   string
	Type string
	Name string
	Args string
}

func (ts *ToolCall) FuncString() (string, error) {
	argList := make(map[string]string)
	err := json.Unmarshal([]byte(ts.Args), &argList)
	if err != nil {
		return "", err
	}
	// Extract the function name and arguments
	functionName := ts.Name
	var args []string
	for _, value := range argList {
		args = append(args, fmt.Sprintf("%q", value))
	}

	// Construct the output string
	return fmt.Sprintf("%s(%s)", functionName, strings.Join(args, ", ")), nil
}

// Generator is an interface for interacting with an AI
type Generator interface {
	Model() string
	Generate(conversation Conversation, meta []Meta, tools []aitool.Tool) (*GeneratorResponse, error)
}
