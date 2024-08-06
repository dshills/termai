package openai

import "github.com/dshills/ai-manager/aitool"

/* -- Request -- */
type CreateRequest struct {
	Model       string            `json:"model,omitempty"`
	Messages    []MessageCallFrag `json:"messages,omitempty"`
	Stream      bool              `json:"stream,omitempty"`
	Logprobs    bool              `json:"logprobs,omitempty"`
	TopLogprobs int               `json:"top_logprobs,omitempty"`
	Tools       []aitool.Tool     `json:"tools,omitempty"`
	ToolChoice  string            `json:"tool_choice,omitempty"`
	Temperature int               `json:"temperature"`
	MaxTokens   *int              `json:"max_tokens,omitempty"`
}

type FunctionFrag struct {
	Name        string         `json:"name,omitempty"`
	Description string         `json:"description,omitempty"`
	Parameters  ParametersFrag `json:"parameters,omitempty"`
}

type ParametersFrag struct {
	Type       string         `json:"type,omitempty"`
	Properties PropertiesFrag `json:"properties,omitempty"`
	Required   []string       `json:"required,omitempty"`
}

type PropertiesFrag struct {
	Location LocationFrag `json:"location,omitempty"`
	Unit     UnitFrag     `json:"unit,omitempty"`
}

type UnitFrag struct {
	Type string   `json:"type,omitempty"`
	Enum []string `json:"enum,omitempty"`
}

type LocationFrag struct {
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
}

type MessageCallFrag struct {
	Role    string `json:"role,omitempty"`
	Content string `json:"content,omitempty"`
}

/* -- Response -- */
type ChatResp struct {
	ID                string       `json:"id,omitempty"`
	Object            string       `json:"object,omitempty"`
	Created           int64        `json:"created,omitempty"`
	Model             string       `json:"model,omitempty"`
	SystemFingerprint string       `json:"system_fingerprint,omitempty"`
	Choices           []ChoiceFrag `json:"choices,omitempty"`
	Usage             UsageFrag    `json:"usage,omitempty"`
}

type ChoiceFrag struct {
	FinishReason string       `json:"finish_reason,omitempty"`
	Index        int64        `json:"index,omitempty"`
	Logprobs     LogprobsFrag `json:"logprobs,omitempty"`
	Message      MessageFrag  `json:"message,omitempty"`
	Delta        MessageFrag  `json:"delta,omitempty"`
}

type UsageFrag struct {
	CompletionTokens int64 `json:"completion_tokens,omitempty"`
	PromptTokens     int64 `json:"prompt_tokens,omitempty"`
	TotalTokens      int64 `json:"total_tokens,omitempty"`
}

type MessageFrag struct {
	Role      string     `json:"role,omitempty"`
	Content   string     `json:"content,omitempty"`
	ToolCalls []ToolFrag `json:"tool_calls,omitempty"`
}

type ToolFrag struct {
	ID       string `json:"id,omitempty"`
	Type     string `json:"type,omitempty"`
	Function struct {
		Name      string `json:"name,omitempty"`
		Arguments string `json:"arguments,omitempty"`
	} `json:"function,omitempty"`
}

type LogprobsFrag struct {
	Content []ContentFrag `json:"content,omitempty"`
}

type ContentFrag struct {
	Token       string           `json:"token,omitempty"`
	Logprob     float64          `json:"logprob,omitempty"`
	Bytes       []int            `json:"bytes,omitempty"`
	TopLogprobs []TopLogprobFrag `json:"top_logprobs,omitempty"`
}

type TopLogprobFrag struct {
	Token   string  `json:"token,omitempty"`
	Logprob float64 `json:"logprob,omitempty"`
	Bytes   []int   `json:"bytes,omitempty"`
}
