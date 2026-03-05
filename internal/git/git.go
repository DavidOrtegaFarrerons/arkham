package git

import (
	"arkham/internal/config"
)

type Git struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Git {
	return &Git{cfg: cfg}
}
func (g *Git) Commit(message string) {
	placeholderValues := g.Parse("branch")
	placeholderValues["message"] = message
	g.Format(placeholderValues)
}
