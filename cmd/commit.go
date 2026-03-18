package cmd

import (
	"arkham/internal/config"
	"arkham/internal/git"
	"errors"

	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:   "commit [message]",
	Short: "Executes a commit based on a branch pattern and a commit template provided by the user",
	Long: `Creates a structured commit message from your current branch and a commit message.

			Arkham reads your current branch name, parses it using your configured branch 
			pattern, and combines it with your commit template to generate a consistent 
			commit message automatically.
			
			Example:
			  Branch:          feature/TASK-1_very-cool-branch
			  Branch pattern:  {type}/{ticket}_{description}
			  Commit template: {type} ({ticket}): {message}
			  Command:         arkham commit "task finished!"
			  Result:          feature (TASK-1): task finished!
			
			Note: {message} is a reserved placeholder that maps to the argument you pass 
			to this command. It cannot be used in your branch pattern.`,

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
