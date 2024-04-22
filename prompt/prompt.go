package prompt

import (
	"fmt"
	"strings"

	"github.com/dshills/ai-manager/ai"
	"github.com/dshills/termai/config"
)

const (
	promptKeyExplain     = "EXPLAIN"
	promptKeyNoExplain   = "NOEXPLAIN"
	promptKeyAIPersona   = "AI-PERSONA"
	promptKeyUserPersona = "USER-PERSONA"
	promptKeyOutput      = "OUTPUT"
	promptKeyOptPrompt   = "OPT-PROMPT"
	replaceFileType      = "%%FILETYPE%%"
)

// Optimize will use the config defined optimizations prompts (OPT_PROMPT)
// to generate a query that ask the AI to optimize the prompt
func Optimize(qry string, prompts []config.Prompt, gen ai.Generator) (string, error) {
	qry = strings.ReplaceAll(qry, "\n", " ")
	opt := extractPromptConfig(prompts, promptKeyOptPrompt, "\n")
	qry = fmt.Sprintf("%s%q", opt, qry)
	conv := ai.Conversation{{Role: "user", Text: qry}}
	resp, err := gen.Generate(conv, nil, nil)
	if err != nil {
		return "", err
	}
	prompt := strings.ReplaceAll(resp.Message.Text, "\n", "")
	prompt = strings.ReplaceAll(prompt, "\t", "")
	prompt = strings.ReplaceAll(prompt, "\"", "")
	return prompt, err
}

// Inject will inject config defined prompt optimizations
// Defaults AI-PERSONA, USER-PERSONA, OUTPUT
// EXPLAIN || NOEXPLAIN
// fileType specific
func Inject(qry, ft string, explain bool, prompts []config.Prompt) string {
	builder := strings.Builder{}

	builder.WriteString(extractPromptConfig(prompts, promptKeyAIPersona, "\n"))
	builder.WriteString(extractPromptConfig(prompts, promptKeyUserPersona, "\n"))
	builder.WriteString(extractPromptConfig(prompts, promptKeyOutput, "\n"))
	if explain {
		builder.WriteString(extractPromptConfig(prompts, promptKeyExplain, "\n"))
	} else {
		builder.WriteString(extractPromptConfig(prompts, promptKeyNoExplain, "\n"))
	}
	builder.WriteString(extractPromptConfig(prompts, ft, "\n"))
	builder.WriteString(qry)
	return strings.ReplaceAll(builder.String(), "%%FILETYPE%%", ft)
}

func extractPromptConfig(prompts []config.Prompt, key, suffix string) string {
	for _, p := range prompts {
		if key == p.Topic {
			return p.Prompt + suffix
		}
	}
	return ""
}
