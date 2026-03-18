package config

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	ErrConfigFileNotFound = errors.New("no config file found")
)

// Config contains BranchPattern and CommitTemplate
// BranchPattern is the pattern of a branch
// eg: for branch: feature/TASK-1_very-cool-branch
// BranchPattern: {type}/{ticket}_{description}
// and CommitTemplate would be a custom pattern that can use all wrapped names in {} in BranchPattern
// eg: {type} ({ticket}): {description}
// This would output: feature (TASK-1): very-cool-branch
// A special pattern "{message}" is allowed, which is the message you input in the commit command
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
	if err != nil {
		return err
	}
	branchPattern = strings.TrimSpace(branchPattern)
	patterns := ExtractPlaceholders(branchPattern)
	patternsMap := map[string]bool{}
	for _, p := range patterns {
		patternsMap[p] = true
	}

	if _, exists := patternsMap["message"]; exists {
		panic("You cannot use a message variable")
	}

	fmt.Println("Add the commit template, you have the following vars available:")
	fmt.Println(strings.Join(patterns, ", "))
	commitTemplate, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	commitTemplate = strings.TrimSpace(commitTemplate)
	commitTemplatePlaceholders := ExtractPlaceholders(commitTemplate)
	for _, templatePlaceholder := range commitTemplatePlaceholders {
		if _, exists := patternsMap[templatePlaceholder]; !exists {
			panic(fmt.Sprintf("%s does not exist as an option based on the branch pattern you wrote", templatePlaceholder))
		}
	}

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

func ExtractPlaceholders(pattern string) []string {
	templateRegex := regexp.MustCompile(`\{(\w+)}`)
	keys := templateRegex.FindAllStringSubmatch(pattern, -1)

	patterns := make([]string, len(keys))
	for i, match := range keys {
		patterns[i] = match[1]
	}

	return patterns
}
