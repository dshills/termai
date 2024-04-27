package main

import (
	"bufio"
	"flag"
	"os"
	"strings"
)

type Options struct {
	ColorOutput    bool
	ExplainOutput  bool
	Help           bool
	InitConfig     bool
	ListModels     bool
	OptimizePrompt bool
	PrintDefModel  bool
	PrintConv      bool
	PrintPrompt    bool
	Continue       bool
	FileType       string
	Model          string
	Out            string
	Query          string
}

func getFlags() Options {
	opts := Options{}
	flag.BoolVar(&opts.ColorOutput, "color", false, "Highlighted output")
	flag.BoolVar(&opts.PrintDefModel, "default", false, "Prints the default model")
	flag.BoolVar(&opts.ExplainOutput, "explain", false, "Explain the solution returned")
	flag.BoolVar(&opts.Help, "help", false, "Print usage information")
	flag.BoolVar(&opts.InitConfig, "init", false, "Generate a default configuration file")
	flag.BoolVar(&opts.ListModels, "list", false, "List available models")
	flag.BoolVar(&opts.PrintPrompt, "prompt", false, "Output the prompt without calling the AI")
	flag.BoolVar(&opts.OptimizePrompt, "opt", false, "Using the selected model try and optimize the prompt")
	flag.BoolVar(&opts.PrintConv, "conv", false, "Print the last conversation")
	flag.BoolVar(&opts.Continue, "continue", false, "Continue last conversation")
	flag.StringVar(&opts.Model, "model", "", "Model to use")
	flag.StringVar(&opts.Out, "out", "", "Output file path")
	flag.StringVar(&opts.FileType, "ft", "", "Use prompt extensions for a specific file type")
	flag.Parse()

	opts.Query = strings.Join(flag.Args(), " ")
	opts.Query += getPipedData()

	return opts
}

func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}

func getPipedData() string {
	if !isInputFromPipe() {
		return ""
	}
	builder := strings.Builder{}
	r := os.Stdin
	scanner := bufio.NewScanner(bufio.NewReader(r))
	for scanner.Scan() {
		builder.WriteString(scanner.Text())
	}
	return builder.String()
}
