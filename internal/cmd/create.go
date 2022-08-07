package cmd

import (
	"github.com/mehditeymorian/gli/internal/builder"
	"github.com/mehditeymorian/gli/internal/config"
	"github.com/mehditeymorian/gli/internal/mod"
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

	cmd.Flags().BoolP("verbose", "v", false, "log verbosity")

	return cmd
}

func run(cmd *cobra.Command, _ []string) {
	cfg := config.Load()

	surveyResult := model.EmptySurveyResult()

	question.New(cfg).Fill(surveyResult)

	app := surveyResult.Execute()
	app.SelectedModules = cfg.GetRequiredModules(surveyResult)
	app.RequiredModules = cfg.RequiredModules

	ExtractFlags(cmd, app)

	builder.NewBuilder(cfg).Build(app)

	mod.DownloadModules(app)
}

func ExtractFlags(cmd *cobra.Command, app *model.App) {
	verbose, err := cmd.Flags().GetBool("verbose")
	if err != nil {
		verbose = false
	}

	app.Flags[model.Verbose] = verbose
}
