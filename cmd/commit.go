package main

import (
	"arkham/internal/config"
	"arkham/internal/git"
	"fmt"

	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:     "commit [message]",
	Short:   "Executes a commit based on some rules",
	Aliases: []string{"c"},
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := config.Config{
			BranchPattern:  "/a/b/c",
			CommitTemplate: "test",
		}
		git := git.New(&cfg)
		git.Commit(args[0])
		fmt.Println("Command executed successfully")
		return nil
	},
}
