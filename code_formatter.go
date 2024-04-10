package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
)

func FormatCodeResponse(text string) string {
	sections := []string{}

	for _, block := range extractCodeBlocks(text) {
		if block.Lang == "" {
			sections = append(sections, block.Content)
			continue
		}
		code, err := highlightCode(block.Content, block.Lang)
		if err != nil {
			fmt.Println(err)
			sections = append(sections, block.Content)
			continue
		}
		sections = append(sections, code)
	}

	return strings.Join(sections, "\n")
}

func highlightCode(sourceCode, language string) (string, error) {
	// Create a new buffer to hold the resulting HTML
	var code bytes.Buffer

	// Determine the lexer based on the language
	lexer := lexers.Get(language)

	// Choose a style for the syntax highlighting.
	style := styles.Get("dracula")
	if style == nil {
		style = styles.Fallback
	}

	// Create an HTML formatter with some options set.
	formatter := formatters.Get("terminal16m")

	// Tokenize the source code.
	iterator, err := lexer.Tokenise(nil, sourceCode)
	if err != nil {
		return "", err
	}

	// Format the tokens into the codefer.
	err = formatter.Format(&code, style, iterator)
	if err != nil {
		return "", err
	}

	return code.String(), nil
}

// Structure to hold extracted code blocks
type CodeBlock struct {
	Lang    string
	Content string
}

func extractCodeBlocks(input string) []CodeBlock {
	lines := strings.Split(input, "\n")
	var codeBlocks []CodeBlock
	var currentBlock CodeBlock

	const prefix = "```"
	inCodeBlock := false

	for _, line := range lines {
		switch {
		case inCodeBlock && strings.HasPrefix(line, prefix):
			// Block end
			inCodeBlock = !inCodeBlock
			if currentBlock.Content != "" {
				codeBlocks = append(codeBlocks, currentBlock)
				currentBlock = CodeBlock{}
			}
		case strings.HasPrefix(line, prefix):
			// Start code block
			inCodeBlock = !inCodeBlock
			if currentBlock.Content != "" {
				codeBlocks = append(codeBlocks, currentBlock)
				currentBlock = CodeBlock{}
			}
			lang := strings.TrimSpace(strings.TrimPrefix(string(line), prefix))
			if lang == "" {
				lang = "unknown"
			}
			currentBlock.Lang = lang
		default:
			currentBlock.Content += line + "\n"
		}
	}

	return codeBlocks
}
