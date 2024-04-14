package config

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/dshills/ai-manager/ai"
	"github.com/dshills/ai-manager/aigen/anthropic"
	"github.com/dshills/ai-manager/aigen/gemini"
	"github.com/dshills/ai-manager/aigen/mistral"
	"github.com/dshills/ai-manager/aigen/ollama"
	"github.com/dshills/ai-manager/aigen/openai"
)

const defPath = "/.termai.json"
const (
	aiAnthropic = "anthropic"
	aiOpenAI    = "openai"
	aiGemini    = "gemini"
	aiOllama    = "ollama"
	aiMistral   = "mistral"
)

type Configuration struct {
	Models       []Model  `json:"models"`
	Prompts      []Prompt `json:"prompts"`
	Generators   []ai.Generator
	ActiveModels []string
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

func LoadDefault() (*Configuration, error) {
	home, err := findHomeDir()
	if err != nil {
		return nil, err
	}

	path := filepath.Join(home, defPath)
	conf, err := LoadConfigPath(path)
	if err != nil {
		return nil, err
	}

	for _, mod := range conf.Models {
		if !mod.Inactive {
			gen, err := makeGenerator(mod)
			if err != nil {
				continue
			}
			conf.Generators = append(conf.Generators, gen)
			conf.ActiveModels = append(conf.ActiveModels, mod.Model)
		}
	}
	return conf, nil
}

func makeGenerator(mod Model) (ai.Generator, error) {
	switch strings.ToLower(mod.Name) {
	case aiOpenAI:
		return openai.New(mod.Model, mod.APIKey, mod.BaseURL), nil
	case aiGemini:
		return gemini.New(mod.Model, mod.APIKey, mod.BaseURL), nil
	case aiOllama:
		return ollama.New(mod.Model, mod.BaseURL), nil
	case aiMistral:
		return mistral.New(mod.Model, mod.APIKey, mod.BaseURL), nil
	case aiAnthropic:
		return anthropic.New(mod.Model, mod.APIKey, mod.BaseURL), nil
	}
	return nil, fmt.Errorf("generator for %v not found", mod.Name)
}
