package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/dshills/ai-manager/ai"
	"github.com/dshills/termai/config"
	"github.com/dshills/termai/format"
	"github.com/dshills/termai/prompt"
	"github.com/dshills/termai/serial"
)

const conversationStore = ".termai.conv"

type AI struct {
	mgr       *ai.Manager
	model     string
	conv      ai.Conversation
	conf      *config.Configuration
	opts      Options
	prompt    string
	orgPrompt string
	defGen    ai.Generator
}

func newAI(generators []ai.Generator, conf *config.Configuration, opts Options) (*AI, error) {
	aim := AI{
		mgr:   ai.New(),
		model: opts.Model,
		conf:  conf,
		opts:  opts,
	}
	// load prev conversation
	_ = aim.loadConv()

	// Set the model
	if aim.model == "" && conf.DefaultModel() == "" {
		return &aim, fmt.Errorf("no model set and no default model")
	}
	if aim.model == "" {
		aim.model = conf.DefaultModel()
	}

	// set default generator
	aim.defGenerator(generators)

	// gather active
	err := aim.mgr.RegisterGenerators(generators...)
	return &aim, err
}

func (a *AI) createPrompt(query string) error {
	a.orgPrompt = query

	// If specifc ft generate an advanced prompt
	if a.opts.FileType != "" {
		query = prompt.Inject(query, a.opts.FileType, a.opts.ExplainOutput, a.conf.Prompts)
	}
	a.prompt = query

	// If opt-prompt add optimization of the prompt to the prompt
	if a.opts.OptimizePrompt {
		optimized, err := prompt.Optimize(query, a.conf.Prompts, a.defGen)
		if err != nil {
			return err
		}
		a.prompt = optimized
	}
	return nil
}

func (a *AI) defGenerator(generators []ai.Generator) {
	for _, g := range generators {
		if strings.EqualFold(g.Model(), a.model) {
			a.defGen = g
			return
		}
	}
}

func (a *AI) printPrompt() {
	fmt.Println(a.prompt)
}

func (a *AI) printDefaultModel() {
	dm := a.conf.DefaultModel()
	if dm == "" {
		fmt.Println("No default model set")
		return
	}
	fmt.Println(dm)
}

func (a *AI) listModels() {
	for _, m := range a.conf.ActiveModels {
		fmt.Println(m)
	}
}

func (a *AI) printConv() {
	for _, msg := range a.conv {
		fmt.Printf("%s: %s\n", msg.Role, msg.Text)
	}
}

func (a *AI) usePrompt() (string, error) {
	// Start a new conversation
	if !a.opts.Continue {
		a.conv = ai.Conversation{}
	}

	// Start the AI
	tData := ai.ThreadData{Model: a.model, Conversation: a.conv}
	thread, err := a.mgr.NewThread(tData)
	if err != nil {
		return "", fmt.Errorf("%s %w", tData.Model, err)
	}

	resp, err := thread.Converse(a.prompt)
	if err != nil {
		return "", fmt.Errorf("%s %w", tData.Model, err)
	}
	a.conv = thread.Conversation()

	if a.opts.ColorOutput {
		fmtOut := format.Response(resp.Message.Text)
		if fmtOut != "" {
			return fmtOut, nil
		}
	}
	return resp.Message.Text, nil
}

func (a *AI) saveConv() error {
	convPath, err := convLocation()
	if err != nil {
		return err
	}
	file, err := os.Create(convPath)
	if err != nil {
		return err
	}
	defer file.Close()
	return serial.Serialize(file, a.conv)
}

func (a *AI) loadConv() error {
	convPath, err := convLocation()
	if err != nil {
		return err
	}
	file, err := os.Open(convPath)
	if err != nil {
		return err
	}
	defer file.Close()
	conv, err := serial.Hydrate(file)
	if err != nil {
		return err
	}
	a.conv = conv
	return nil
}

func findHomeDir() (string, error) {
	home := os.Getenv("HOME")
	if home == "" {
		return "", fmt.Errorf("could not determine HOME directory")
	}
	return home, nil
}

func convLocation() (string, error) {
	home, err := findHomeDir()
	if err != nil {
		return "", err
	}
	fn := filepath.Join(home, conversationStore)
	fn, err = filepath.Abs(fn)
	if err != nil {
		return "", err
	}
	return fn, nil
}
