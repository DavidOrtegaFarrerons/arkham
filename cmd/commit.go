package main

import (
	"arkham/internal/config"
	"arkham/internal/git"
	"errors"
	"fmt"

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
			//Ask for info
		}
		g := git.New(cfg)
		g.Commit(args[0])
		fmt.Println("Command executed successfully")
		return nil
	},
}
