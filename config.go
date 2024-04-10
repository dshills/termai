package main

import (
	"fmt"
	"strings"
)

const (
	promptAll       = "ALL"
	promptExplain   = "EXPLAIN"
	promptNoExplain = "NOEXPLAIN"
	replaceFileType = "%%FILETYPE%%"
)

type Configuration struct {
	Models  []Model  `json:"models"`
	Prompts []Prompt `json:"prompts"`
}

type Model struct {
	Name     string `json:"name"`
	Model    string `json:"model"`
	APIKey   string `json:"api_key"`
	BaseURL  string `json:"base_url"`
	Inactive bool   `json:"inactive"`
	Default  bool   `json:"default"`
}

type Prompt struct {
	Topic  string `json:"topic"`
	Prompt string `json:"prompt"`
}

func (c Configuration) Prompt(query, fileType string, shouldExplain bool) string {
	return fmt.Sprintf("%s%s%s%s", c.all(fileType), c.explain(fileType, shouldExplain), c.extra(fileType), query)
}

func (c Configuration) extra(ft string) string {
	ft = strings.ToLower(ft)
	for _, p := range c.Prompts {
		if strings.ToLower(p.Topic) == ft {
			return strings.ReplaceAll(p.Prompt, replaceFileType, ft) + "\n"
		}
	}
	return ""
}

func (c Configuration) all(ft string) string {
	for _, p := range c.Prompts {
		if p.Topic == promptAll {
			return strings.ReplaceAll(p.Prompt, replaceFileType, ft) + "\n"
		}
	}
	return ""
}

func (c Configuration) explain(ft string, shouldExplain bool) string {
	for _, p := range c.Prompts {
		if shouldExplain && p.Topic == promptExplain {
			return strings.ReplaceAll(p.Prompt, replaceFileType, ft) + "\n"
		}
		if !shouldExplain && p.Topic == promptNoExplain {
			return strings.ReplaceAll(p.Prompt, replaceFileType, ft) + "\n"
		}
	}
	return ""
}

func (c Configuration) DefaultModel() Model {
	for _, m := range c.Models {
		if m.Default {
			return m
		}
	}
	return Model{}
}
