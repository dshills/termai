package serial

import (
	"encoding/gob"
	"io"

	"github.com/dshills/ai-manager/ai"
)

type What int

const (
	Response What = 1
	Prompt   What = 2
)

type Where int

const (
	First Where = 1
	Last  Where = 2
)

func Hydrate(r io.Reader) (ai.Conversation, error) {
	conv := ai.Conversation{}
	err := gob.NewDecoder(r).Decode(&conv)
	return conv, err
}

func Extract(what What, where Where, conv ai.Conversation) string {
	if len(conv) == 0 {
		return ""
	}

	if where == First {
		for _, msg := range conv {
			if what == Prompt && msg.Role == "user" {
				return msg.Text
			}
			if what == Response && msg.Role == "assistant" {
				return msg.Text
			}
		}
		return ""
	}

	if where == Last {
		for i := len(conv) - 1; i >= 0; i-- {
			if what == Prompt && conv[i].Role == "user" {
				return conv[i].Text
			}
			if what == Response && conv[i].Role == "assistant" {
				return conv[i].Text
			}
		}
		return ""
	}

	return ""
}
