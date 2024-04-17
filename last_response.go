package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const (
	lastRespPath   = "/.termai.last"
	lastPromptPath = "/.termai.last.prompt"
)

func findHomeDir() (string, error) {
	home := os.Getenv("HOME")
	if home == "" {
		return "", fmt.Errorf("could not determine HOME directory")
	}
	return home, nil
}

func saveResponse(response string) error {
	home, err := findHomeDir()
	if err != nil {
		return fmt.Errorf("saveResponse: %w", err)
	}
	fname := filepath.Join(home, lastRespPath)
	file, err := os.Create(fname)
	if err != nil {
		return fmt.Errorf("saveResponse: Create: %w", err)
	}
	defer file.Close()
	if _, err := file.WriteString(response); err != nil {
		return fmt.Errorf("saveResponse: WriteString: %w", err)
	}
	return nil
}

func savePrompt(prompt string) error {
	home, err := findHomeDir()
	if err != nil {
		return fmt.Errorf("savePrompt: %w", err)
	}
	fname := filepath.Join(home, lastPromptPath)
	file, err := os.Create(fname)
	if err != nil {
		return fmt.Errorf("savePrompt: Create: %w", err)
	}
	defer file.Close()
	if _, err := file.WriteString(prompt); err != nil {
		return fmt.Errorf("savePrompt: WriteString: %w", err)
	}
	return nil
}

func lastResponse() (string, error) {
	home, err := findHomeDir()
	if err != nil {
		return "", fmt.Errorf("lastResponse: %w", err)
	}
	fname := filepath.Join(home, lastRespPath)
	file, err := os.Open(fname)
	if err != nil {
		return "", fmt.Errorf("lastResponse: Create: %w", err)
	}
	defer file.Close()
	byts, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("lastResponse: Create: %w", err)
	}
	return string(byts), nil
}

func lastPrompt() (string, error) {
	home, err := findHomeDir()
	if err != nil {
		return "", fmt.Errorf("lastPrompt: %w", err)
	}
	fname := filepath.Join(home, lastPromptPath)
	file, err := os.Open(fname)
	if err != nil {
		return "", fmt.Errorf("lastPrompt: Create: %w", err)
	}
	defer file.Close()
	byts, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("lastPrompt: Create: %w", err)
	}
	return string(byts), nil
}
