package main

import (
	"fmt"
	"strings"

	"github.com/dshills/ai-manager/ai"
)

type Configuration struct {
	Models   []Model  `json:"models"`
	Prompts  []Prompt `json:"prompts"`
	aiModels []ai.Model
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

func (c Configuration) DefaultModel() string {
	for _, m := range c.Models {
		if m.Default {
			return m.Model
		}
	}
	return ""
}

func (c Configuration) AIModels() []ai.Model {
	return c.aiModels
}

func (c Configuration) GetAIModel(modName string) (*ai.Model, error) {
	modName = strings.ToLower(modName)
	for _, mod := range c.aiModels {
		if modName == strings.ToLower(mod.Model) {
			return &mod, nil
		}
	}
	return nil, fmt.Errorf("model not found")
}

func (c Configuration) ListModels() []string {
	models := []string{}
	for _, mod := range c.aiModels {
		models = append(models, mod.Model)
	}
	return models
}
