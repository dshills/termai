package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/dshills/ai-manager/ai"
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
	defaults := false
	explain := false
	prompt := false
	init := false
	list := false
	help := false
	aiModel := ""
	color := false
	flag.StringVar(&fileType, "ft", fileType, "Use prompt extensions for a specific file type")
	flag.BoolVar(&prompt, "prompt", prompt, "Output the prompt without calling the AI")
	flag.BoolVar(&explain, "explain", explain, "Explain the solution returned")
	flag.BoolVar(&init, "init", init, "Generate a default configuration file")
	flag.BoolVar(&list, "list", list, "List available models")
	flag.BoolVar(&defaults, "defaults", defaults, "Prints the default model")
	flag.BoolVar(&help, "help", help, "Print usage information")
	flag.BoolVar(&color, "color", color, "Highlighted output")
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
	query := strings.Join(flag.Args(), " ")

	conf, err := LoadDefaultConfig()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Usage: termai [options] [query]")
		flag.PrintDefaults()
		os.Exit(1)
	}

	defModel := conf.DefaultModel()

	// Defaults: Print out defaults
	if defaults {
		if defModel == "" {
			fmt.Println("No default set")
		} else {
			fmt.Println(defModel)
		}
		os.Exit(0)
	}
	// List:  List Models
	if list {
		for _, m := range conf.ListModels() {
			fmt.Println(m)
		}
		os.Exit(0)
	}
	// Help: Print usage
	if help {
		fmt.Println("Usage: termai [options] [query]")
		flag.PrintDefaults()
		os.Exit(0)
	}

	// At this point we need a query
	if query == "" {
		fmt.Println("You didn't ask anything")
		fmt.Println("Usage: termai [options] [query]")
		flag.PrintDefaults()
		os.Exit(1)
	}
	// If specifc ft generate an advanced prompt
	if fileType != "" {
		query = conf.Prompt(query, fileType, explain)
	}
	// Prompt: Print the prompt without running
	if prompt {
		fmt.Println(query)
		os.Exit(0)
	}

	if aiModel == "" {
		aiModel = defModel
	}
	if aiModel == "" {
		fmt.Println("No default model set")
		fmt.Println("Usage: termai [options] [query]")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Start the AI
	aimgr := ai.New()

	if err := aimgr.RegisterGenerators(conf.AIModels()...); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	aimod, err := conf.GetAIModel(aiModel)
	if err != nil {
		fmt.Printf("model %v not found\n", aiModel)
		os.Exit(1)
	}

	tData := ai.ThreadData{AIName: aimod.AIName, Model: aimod.Model}
	thread, err := aimgr.NewThread(tData)
	if err != nil {
		fmt.Printf("%s %s %v\n", aimod.AIName, aimod.Model, err)
		os.Exit(1)
	}

	resp, err := thread.Converse(query)
	if err != nil {
		fmt.Printf("%s %s %v\n", aimod.AIName, aimod.Model, err)
		os.Exit(1)
	}

	if color {
		fmtOut := FormatCodeResponse(resp.Message.Text)
		if fmtOut != "" {
			fmt.Println(fmtOut)
			os.Exit(0)
		}
	}
	fmt.Println(resp.Message.Text)
}
