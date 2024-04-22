package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dshills/termai/config"
)

func main() {
	opts := getFlags()

	// Create defaults
	if opts.InitConfig {
		if err := config.InitializeDefaults(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	// load config
	conf, err := config.LoadDefault()
	if err != nil {
		ShowUsageAndExit(err.Error(), 1)
	}

	// Create the ai
	aimgr, err := newAI(conf.Generators, conf, opts)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Print the conversation
	if opts.PrintConv {
		aimgr.printConv()
		os.Exit(0)
	}

	// Print the default model
	if opts.PrintDefModel {
		aimgr.printDefaultModel()
		os.Exit(0)
	}

	// List:  List Models
	if opts.ListModels {
		aimgr.listModels()
		os.Exit(0)
	}

	// Help: Print usage
	if opts.Help {
		ShowUsageAndExit("", 0)
	}

	// --- Query ---
	query := opts.Query
	if query == "" {
		ShowUsageAndExit("You didn't ask anything", 1)
	}

	if err := aimgr.createPrompt(query); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Prompt: Print the prompt without running
	if opts.PrintPrompt {
		aimgr.printPrompt()
		os.Exit(0)
	}

	output, err := aimgr.usePrompt()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(output)

	if err := aimgr.saveConv(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ShowUsageAndExit(msg string, exitcode int) {
	if msg != "" {
		fmt.Println(msg)
	}
	fmt.Println("Usage: termai [options] [query]")
	flag.PrintDefaults()
	os.Exit(exitcode)

}
