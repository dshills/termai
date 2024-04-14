package config

import "fmt"

func InitializeDefaults() error {
	if err := writeDefConfig(); err != nil {
		return err
	}

	fmt.Println("$HOME/.termai.json configuration file created.")
	fmt.Println("1) Open the file")
	fmt.Println("2) Add your API keys")
	fmt.Println("3) Mark models you wish to use as Active")
	fmt.Println("4) Mark one model as Default (Can be overridden)")
	fmt.Println("5) Add any langugae specfic prompts to the \"prompts\" section")
	return nil
}

func (c Configuration) DefaultModel() string {
	for _, m := range c.Models {
		if m.Default {
			return m.Model
		}
	}
	return ""
}

var defaultConfig = Configuration{
	Models: []Model{
		{
			Name:     "Gemini",
			Model:    "gemini-pro",
			APIKey:   "",
			BaseURL:  "https://generativelanguage.googleapis.com/v1beta",
			Inactive: true,
			Default:  false,
		},
		{
			Name:     "Gemini",
			Model:    "gemini-1.0-pro-latest",
			APIKey:   "",
			BaseURL:  "https://generativelanguage.googleapis.com/v1beta",
			Inactive: true,
			Default:  false,
		},
		{
			Name:     "Gemini",
			Model:    "gemini-1.0-pro",
			APIKey:   "",
			BaseURL:  "https://generativelanguage.googleapis.com/v1beta",
			Inactive: true,
			Default:  false,
		},
		{
			Name:     "OpenAI",
			Model:    "gpt-4",
			APIKey:   "",
			BaseURL:  "https://api.openai.com/v1",
			Inactive: true,
			Default:  false,
		},
		{
			Name:     "OpenAI",
			Model:    "gpt-3.5-turbo",
			APIKey:   "",
			BaseURL:  "https://api.openai.com/v1",
			Inactive: true,
			Default:  false,
		},
		{
			Name:     "OpenAI",
			Model:    "gpt-4-turbo-preview",
			APIKey:   "",
			BaseURL:  "https://api.openai.com/v1",
			Inactive: true,
			Default:  false,
		},
		{
			Name:     "Mistral",
			Model:    "mistral-small-latest",
			APIKey:   "",
			BaseURL:  "https://api.mistral.ai/v1",
			Inactive: true,
			Default:  false,
		},
		{
			Name:     "Mistral",
			Model:    "mistral-medium-latest",
			APIKey:   "",
			BaseURL:  "https://api.mistral.ai/v1",
			Inactive: true,
			Default:  false,
		},
		{
			Name:     "Mistral",
			Model:    "mistral-large-latest",
			APIKey:   "",
			BaseURL:  "https://api.mistral.ai/v1",
			Inactive: true,
			Default:  false,
		},
		{
			Name:     "Anthropic",
			Model:    "claude-3-opus-20240229",
			APIKey:   "",
			BaseURL:  "https://api.anthropic.com/v1",
			Inactive: true,
			Default:  false,
		},
		{
			Name:     "Anthropic",
			Model:    "claude-3-sonnet-20240229",
			APIKey:   "",
			BaseURL:  "https://api.anthropic.com/v1",
			Inactive: true,
			Default:  false,
		},
		{
			Name:     "Anthropic",
			Model:    "claude-3-haiku-20240307",
			APIKey:   "",
			BaseURL:  "https://api.anthropic.com/v1",
			Inactive: true,
			Default:  false,
		},
		{
			Name:     "Ollama",
			Model:    "starcoder2:15b",
			BaseURL:  "http://localhost:11434",
			Inactive: true,
			Default:  false,
		},
	},
	Prompts: []Prompt{
		{
			Topic:  "OPT-PROMPT",
			Prompt: "You are an expert in prompt engineering.\nRewrite this AI prompt to get the best results for code generation.\nThe text appearing inside of quotes is the prompt to be optimized.",
		},
		{
			Topic:  "AI-PERSONA",
			Prompt: "Act as a highly experienced software developer specializing in %%FILETYPE%%",
		},
		{
			Topic:  "USER-PERSONA",
			Prompt: "Explain it to a highly experienmced %%FILETYPE%% developer.",
		},
		{
			Topic:  "OUTPUT",
			Prompt: "Your work should be expertly written with unique code comments for all functions and data structures. Your task is to create fully functional and bug free code.",
		},
		{
			Topic:  "EXPLAIN",
			Prompt: "Explain, in detail, the returned code. Explain why you made the choices you did and why I would want to do it this way.",
		},
		{
			Topic:  "NOEXPLAIN",
			Prompt: "Provide only code with comments and no explanations.",
		},
		{Topic: "go", Prompt: ""},
		{Topic: "react", Prompt: ""},
		{Topic: "ts", Prompt: ""},
	},
}
