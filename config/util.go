package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

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

	return &conf, nil
}

func writeDefConfig() error {
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
