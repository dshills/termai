package main

import (
	"encoding/json"
	"fmt"
	"os"
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

func findHomeDir() (string, error) {
	home := os.Getenv("HOME")
	if home == "" {
		return "", fmt.Errorf("could not determine HOME directory")
	}
	return home, nil
}

func fileExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func LoadConfigPath(fpath string) (*Configuration, error) {
	file, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	conf := Configuration{}
	if err := json.NewDecoder(file).Decode(&conf); err != nil {
		return nil, err
	}

	for _, m := range conf.Models {
		mod, err := convertModel(m)
		if err != nil || mod == nil {
			continue
		}
		conf.aiModels = append(conf.aiModels, *mod)
	}
	return &conf, nil
}

func LoadDefaultConfig() (*Configuration, error) {
	home, err := findHomeDir()
	if err != nil {
		return nil, err
	}

	path := filepath.Join(home, defPath)
	return LoadConfigPath(path)
}

func InitializeDefConfig() error {
	home, err := findHomeDir()
	if err != nil {
		return err
	}
	path := filepath.Join(home, defPath)

	if fileExists(path) {
		return fmt.Errorf("config file exists %v", path)
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")
	return encoder.Encode(defaultConfig)
}

func makeGenerator(aiName string) (ai.Generator, error) {
	switch strings.ToLower(aiName) {
	case aiOpenAI:
		return openai.New(), nil
	case aiGemini:
		return gemini.New(), nil
	case aiOllama:
		return ollama.New(), nil
	case aiMistral:
		return mistral.New(), nil
	case aiAnthropic:
		return anthropic.New(), nil
	}
	return nil, fmt.Errorf("generator for %v not found", aiName)
}

func convertModel(model Model) (*ai.Model, error) {
	if model.Inactive {
		return nil, nil
	}
	gen, err := makeGenerator(model.Name)
	if err != nil {
		return nil, err
	}
	return &ai.Model{
		AIName:    model.Name,
		Model:     model.Model,
		APIKey:    model.APIKey,
		BaseURL:   model.BaseURL,
		Generator: gen,
	}, nil
}
