package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dshills/ai-manager/ai"
	"github.com/dshills/termai/config"
	"github.com/dshills/termai/prompt"
)

func main() {
	opts := handleFlags()

	if opts.InitConfig {
		if err := config.InitializeDefaults(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	opts.Query += getPipedData()

	conf, err := config.LoadDefault()
	if err != nil {
		ShowUsageAndExit(err.Error(), 1)
	}

	aimgr := ai.New()
	if err := aimgr.RegisterGenerators(conf.Generators...); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defModel := conf.DefaultModel()

	// Defaults: Print out defaults
	if opts.ShowDefaults {
		if defModel == "" {
			fmt.Println("No default set")
		} else {
			fmt.Println(defModel)
		}
		os.Exit(0)
	}
	// List:  List Models
	if opts.ListModels {
		for _, m := range conf.ActiveModels {
			fmt.Println(m)
		}
		os.Exit(0)
	}
	// Help: Print usage
	if opts.Help {
		ShowUsageAndExit("", 0)
	}

	// At this point we need a query
	if opts.Query == "" {
		ShowUsageAndExit("You didn't ask anything", 1)
	}
	// If specifc ft generate an advanced prompt
	if opts.FileType != "" {
		opts.Query = prompt.Inject(opts.Query, opts.FileType, opts.ExplainOutput, conf.Prompts)
	}

	// If opt-prompt add optimization of the prompt to the prompt
	if opts.OptimizePrompt || opts.UseOptPrompt {
		opts.Query = prompt.Optimize(opts.Query, conf.Prompts)
	}
	// Prompt: Print the prompt without running
	if opts.ShowPrompt {
		fmt.Println(opts.Query)
		os.Exit(0)
	}

	if opts.Model == "" {
		opts.Model = defModel
	}
	if opts.Model == "" {
		ShowUsageAndExit("No model set", 1)
	}

	// Start the AI
	tData := ai.ThreadData{Model: opts.Model}
	resp := converse(aimgr, tData, opts.Query)

	// if opt-prompt-send optimized prompt and use it
	if opts.UseOptPrompt {
		fmt.Println("Optimized Prompt: " + resp.Message.Text)
		tData := ai.ThreadData{Model: opts.Model}
		resp = converse(aimgr, tData, resp.Message.Text)
	}

	if opts.ColorOutput {
		fmtOut := FormatCodeResponse(resp.Message.Text)
		if fmtOut != "" {
			fmt.Println(fmtOut)
			os.Exit(0)
		}
	}
	fmt.Println(resp.Message.Text)
}

func ShowUsageAndExit(msg string, exitcode int) {
	if msg != "" {
		fmt.Println(msg)
	}
	fmt.Println("Usage: termai [options] [query]")
	flag.PrintDefaults()
	os.Exit(exitcode)

}
