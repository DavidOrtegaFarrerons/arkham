package cmd

import (
	"arkham/internal/config"
	"arkham/internal/git"
	"errors"

	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:     "commit [message]",
	Short:   "Executes a commit based on some rules",
	Aliases: []string{"c"},
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if errors.Is(err, config.ErrConfigFileNotFound) {
			//ASK INFO
			err := config.Prompt()
			if err != nil {
				panic(err)
			}

			cfg, err = config.Load()
			if err != nil {
				panic(err)
			}
		}
		g := git.New(cfg)
		g.Commit(args[0])
		return nil
	},
}
