package cmd

import (
	"github.com/mehditeymorian/gli/internal/builder"
	"github.com/mehditeymorian/gli/internal/config"
	"github.com/mehditeymorian/gli/internal/model"
	"github.com/mehditeymorian/gli/internal/question"
	"github.com/spf13/cobra"
)

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
	cfg := config.Load()

	app := model.EmptyApp()

	question.New(cfg).Fill(app)

	app.Execute()

	builder.NewBuilder(cfg).Build(app)
}
