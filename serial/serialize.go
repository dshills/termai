package serial

import (
	"encoding/gob"
	"io"

	"github.com/dshills/ai-manager/ai"
)

func Serialize(w io.Writer, conv ai.Conversation) error {
	return gob.NewEncoder(w).Encode(&conv)
}
