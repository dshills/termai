package main

import (
	"fmt"
	"strings"
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

func optimizePrompt(qry string, prompts []Prompt) string {
	qry = strings.ReplaceAll(qry, "\n", " ")
	opt := extractPromptConfig(prompts, promptKeyOptPrompt, "\n")
	return fmt.Sprintf("%s%q", opt, qry)
}

func promptInject(qry, ft string, explain bool, prompts []Prompt) string {
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

func extractPromptConfig(prompts []Prompt, key, suffix string) string {
	for _, p := range prompts {
		if key == p.Topic {
			return p.Prompt + suffix
		}
	}
	return ""
}
