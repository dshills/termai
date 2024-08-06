package ollama

import (
	"time"

	"github.com/dshills/ai-manager/ai"
)

const AIName = "ollama"

const (
	roleAssistant = "assistant"
	roleUser      = "user"
)

type Options struct {
	NumKeep            int      `json:"num_keep,omitempty"`
	Seed               int      `json:"seed,omitempty"`
	NumPredict         int      `json:"num_predict,omitempty"`
	TopK               int      `json:"top_k,omitempty"`
	TopP               float64  `json:"top_p,omitempty"`
	TfsZ               float64  `json:"tfs_z,omitempty"`
	TypicalP           float64  `json:"typical_p,omitempty"`
	RepeatLastN        int      `json:"repeat_last_n,omitempty"`
	Temperature        float64  `json:"temperature,omitempty"`
	RepeatPenalty      float64  `json:"repeat_penalty,omitempty"`
	PresencePenalty    float64  `json:"presence_penalty,omitempty"`
	FrequencyPenalty   float64  `json:"frequency_penalty,omitempty"`
	Mirostat           int      `json:"mirostat,omitempty"`
	MirostatTau        float64  `json:"mirostat_tau,omitempty"`
	MirostatEta        float64  `json:"mirostat_eta,omitempty"`
	PenalizeNewline    bool     `json:"penalize_newline,omitempty"`
	Stop               []string `json:"stop,omitempty"`
	Numa               bool     `json:"numa,omitempty"`
	NumCtx             int      `json:"num_ctx,omitempty"`
	NumBatch           int      `json:"num_batch,omitempty"`
	NumGqa             int      `json:"num_gqa,omitempty"`
	NumGpu             int      `json:"num_gpu,omitempty"`
	MainGpu            int      `json:"main_gpu,omitempty"`
	LowVram            bool     `json:"low_vram,omitempty"`
	F16Kv              bool     `json:"f16_kv,omitempty"`
	VocabOnly          bool     `json:"vocab_only,omitempty"`
	UseMmap            bool     `json:"use_mmap,omitempty"`
	UseMlock           bool     `json:"use_mlock,omitempty"`
	RopeFrequencyBase  float64  `json:"rope_frequency_base,omitempty"`
	RopeFrequencyScale float64  `json:"rope_frequency_scale,omitempty"`
	NumThread          int      `json:"num_thread,omitempty"`
}

type ChatRequest struct {
	Model    string        `json:"model"`
	Messages []MessageFrag `json:"messages"`
	Stream   bool          `json:"stream"`
	Options  Options       `json:"options"`
}

func (cr *ChatRequest) convConv(conversation ai.Conversation) {
	for _, c := range conversation {
		cr.Messages = append(cr.Messages, MessageFrag{Role: c.Role, Content: c.Text})
	}
}

type MessageFrag struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatResponse struct {
	Model     string    `json:"model"`
	CreatedAt time.Time `json:"created_at"`
	Message   struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"message"`
	Done               bool  `json:"done"`
	TotalDuration      int64 `json:"total_duration"`
	LoadDuration       int   `json:"load_duration"`
	PromptEvalCount    int   `json:"prompt_eval_count"`
	PromptEvalDuration int   `json:"prompt_eval_duration"`
	EvalCount          int   `json:"eval_count"`
	EvalDuration       int64 `json:"eval_duration"`
}
