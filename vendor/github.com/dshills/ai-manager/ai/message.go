package ai

// Conversation is a running dialog with an LLM
type Conversation []Message

// Meta is used for storing and passing other pieces of data
type Meta struct {
	Key   string
	Value string
}

// Message represents a single interaction from an LLM
// Each LLM has it's own format for interactions but
// each one has some concept of role and text
type Message struct {
	Role string
	Text string
}

type Usage struct {
	CompletionTokens int64
	PromptTokens     int64
	TotalTokens      int64
}
