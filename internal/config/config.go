package config

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var (
	ErrConfigFileNotFound = errors.New("no config file found")
)

type Config struct {
	BranchPattern  string `json:"branch_pattern"`
	CommitTemplate string `json:"commit_template"`
}

func Load() (*Config, error) {
	path, err := configPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, ErrConfigFileNotFound
	}

	var cfg Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func Save(cfg *Config) error {
	path, err := configPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func Prompt() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Add the branch pattern, eg: {type}/{ticket}_{description}")
	branchPattern, err := reader.ReadString('\n')
	//Here I would like to check if any pattern is "message" and return an error
	if err != nil {
		return err
	}

	//Here I would like to show all the available patterns + "message"
	fmt.Println("Add the commit template, you have the following vars available:")
	fmt.Println()
	commitTemplate, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	//Check here if any pattern from commitTemplate does not exist from branchPattern and return an error

	cfg := &Config{
		BranchPattern:  branchPattern,
		CommitTemplate: commitTemplate,
	}

	return Save(cfg)
}

func configPath() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	appDir := filepath.Join(dir, "arkham")

	if err := os.MkdirAll(appDir, 0755); err != nil {
		return "", err
	}

	return filepath.Join(appDir, "config.json"), nil
}
