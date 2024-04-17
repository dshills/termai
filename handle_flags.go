package main

import (
	"flag"
	"strings"
)

type Options struct {
	ColorOutput    bool
	ExplainOutput  bool
	Help           bool
	InitConfig     bool
	ListModels     bool
	OptimizePrompt bool
	ShowDefaults   bool
	ShowPrompt     bool
	UseOptPrompt   bool
	ShowLast       bool
	ShowLastPrompt bool
	UseLast        bool
	FileType       string
	Model          string
	Query          string
}

func handleFlags() Options {
	opts := Options{}
	flag.BoolVar(&opts.ColorOutput, "color", false, "Highlighted output")
	flag.BoolVar(&opts.ShowDefaults, "defaults", false, "Prints the default model")
	flag.BoolVar(&opts.ExplainOutput, "explain", false, "Explain the solution returned")
	flag.BoolVar(&opts.Help, "help", false, "Print usage information")
	flag.BoolVar(&opts.InitConfig, "init", false, "Generate a default configuration file")
	flag.BoolVar(&opts.ListModels, "list", false, "List available models")
	flag.BoolVar(&opts.ShowPrompt, "prompt", false, "Output the prompt without calling the AI")
	flag.BoolVar(&opts.OptimizePrompt, "opt-prompt", false, "Using the selected model try and optimize the prompt")
	flag.BoolVar(&opts.UseOptPrompt, "opt-prompt-send", false, "Optimize the prompt and then use it")
	flag.BoolVar(&opts.ShowLast, "last", false, "Show last response ($HOME/.termai.last)")
	flag.BoolVar(&opts.UseLast, "use-last", false, "Include the last prompt and response in query")
	flag.BoolVar(&opts.ShowLastPrompt, "last-prompt", false, "Show last prompt used ($HOME/.termai.last.prompt)")
	flag.StringVar(&opts.Model, "model", "", "Model to use")
	flag.StringVar(&opts.FileType, "ft", "", "Use prompt extensions for a specific file type")
	flag.Parse()

	opts.Query = strings.Join(flag.Args(), " ")

	return opts
}
