package git

import (
	"arkham/internal/config"
	"fmt"
	"os/exec"
)

type Git struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Git {
	return &Git{cfg: cfg}
}

func (g *Git) currentBranch() (string, error) {
	o, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		return "", err
	}

	return string(o), nil
}
func (g *Git) Commit(message string) {
	currentBranch, err := g.currentBranch()
	if err != nil {
		panic(err)
	}

	placeholderValues := g.Parse(currentBranch)
	placeholderValues["message"] = message
	commitMsg := g.Format(placeholderValues)

	o, err := exec.Command("git", "commit", "-m", commitMsg).Output()
	if err != nil {
		panic(err)
	}

	fmt.Println(o)
}
