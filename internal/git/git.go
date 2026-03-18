package git

import (
	"arkham/internal/config"
	"os/exec"
)

type Git struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Git {
	return &Git{cfg: cfg}
}
func (g *Git) Commit(message string) {
	placeholderValues := g.Parse("feature/TASK-1_example_for_testing")
	placeholderValues["message"] = message
	commitMsg := g.Format(placeholderValues)

	exec.Command("git", "commit", "-m", commitMsg)
}
