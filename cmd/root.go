package main

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:     "arkham",
	Short:   "A CLI tool for everything you have ever dreamed of",
	Aliases: []string{"ark"},
}

func Execute() {
	rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(commitCmd)
}
