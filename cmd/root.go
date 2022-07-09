package cmd

import (
	"github.com/mehditeymorian/gli/internal/cmd"
	"github.com/spf13/cobra"
	"log"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:   "gli",
		Short: "Go CLI for generating boilerplate code",
		Long:  `Generate boilerplate go code for your project.`,
	}

	rootCmd.AddCommand(
		cmd.Create(),
	)

	if err := rootCmd.Execute(); err != nil {
		log.Panicln(err)
	}
}
