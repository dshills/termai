package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/dshills/ai-manager/ai"
	"github.com/dshills/ai-manager/aigen/anthropic"
	"github.com/dshills/ai-manager/aigen/gemini"
	"github.com/dshills/ai-manager/aigen/mistral"
	"github.com/dshills/ai-manager/aigen/ollama"
	"github.com/dshills/ai-manager/aigen/openai"
)

const (
	aiAnthropic = "anthropic"
	aiOpenAI    = "openai"
	aiGemini    = "gemini"
	aiOllama    = "ollama"
	aiMistral   = "mistral"
)

func main() {
	fileType := ""
	explain := false
	dryRun := false
	init := false
	aiName := ""
	aiModel := ""
	flag.StringVar(&fileType, "ft", fileType, "Use prompt extensions for a specific file type")
	flag.BoolVar(&dryRun, "dryrun", dryRun, "Output the prompt without calling the AI")
	flag.BoolVar(&explain, "explain", explain, "Explain the solution returned")
	flag.BoolVar(&init, "init", init, "Generate a default configuration file")

	conf, err := LoadDefault()
	if err != nil {
		flag.StringVar(&aiName, "ai", aiName, "AI to use")
		flag.StringVar(&aiModel, "model", aiModel, "Model to use")
		flag.Parse()
		if init {
			if err := InitializeDefConfig(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println("$HOME/.termai.json configuration file created.")
			fmt.Println("1) Open the file")
			fmt.Println("2) Add your API keys")
			fmt.Println("3) Mark models you wish to use as Active")
			fmt.Println("4) Mark one model as Default (Can be overridden)")
			fmt.Println("5) Add any langugae specfic prompts to the \"prompts\" section")
			os.Exit(0)
		}
		fmt.Println(err)
		flag.PrintDefaults()
		os.Exit(1)
	}

	defModel := conf.DefaultModel()
	aiName = defModel.Name
	aiModel = defModel.Model

	flag.StringVar(&aiName, "ai", aiName, "AI to use")
	flag.StringVar(&aiModel, "model", aiModel, "Model to use")
	flag.Parse()

	if init {
		fmt.Println("Configuration already exists")
		flag.PrintDefaults()
		os.Exit(1)
	}

	query := strings.Join(flag.Args(), " ")

	if aiName == "" || aiModel == "" {
		fmt.Println("No default model set")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if query == "" {
		fmt.Println("You didn't ask anything")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if fileType != "" {
		query = conf.Prompt(query, fileType, explain)
	}

	if dryRun {
		fmt.Println(query)
		os.Exit(0)
	}

	aimgr := ai.New()

	models := makeModels(conf)
	if err := aimgr.RegisterGenerators(models...); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tData := ai.ThreadData{AIName: aiName, Model: aiModel}
	thread, err := aimgr.NewThread(tData)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	resp, err := thread.Converse(query)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(resp.Message.Text)
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

func makeModels(conf *Configuration) []ai.Model {
	models := []ai.Model{}
	for _, m := range conf.Models {
		gen, err := makeGenerator(m.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		mod := ai.Model{AIName: m.Name, Model: m.Model, APIKey: m.APIKey, BaseURL: m.BaseURL, Generator: gen}
		models = append(models, mod)
	}
	return models
}
