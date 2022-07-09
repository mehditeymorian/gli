package cmd

import "github.com/spf13/cobra"

func Create() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new project",
		Long:  `Create a new project.`,
		Run:   run,
	}

	return cmd
}

func run(_ *cobra.Command, _ []string) {
	
}