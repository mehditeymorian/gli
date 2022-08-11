package question

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/mehditeymorian/gli/internal/config"
	"github.com/mehditeymorian/gli/internal/model"
)

type Question struct {
	Config config.Config
}

func New(cfg config.Config) Question {
	return Question{
		Config: cfg,
	}
}

func (q Question) Fill(app *model.SurveyResult) {
	err := survey.Ask(InitialQuestions(q.Config), app)
	handleErr("failed to ask general questions", err)

	err = survey.Ask(ModuleQuestions(app.Modules, q.Config), app)
	handleErr("failed to ask modules", err)

	err = survey.AskOne(StartPointQuestion(), &app.CliType)
	handleErr("failed to ask start point", err)

	if app.CliType {
		err = survey.AskOne(CliNameQuestion(), &app.CliName)
		handleErr("failed to ask cli name", err)
	}

	err = survey.AskOne(DockerfileQuestion(), &app.Dockerfile)
	handleErr("failed to ask dockerfile", err)
}

func InitialQuestions(cfg config.Config) []*survey.Question {
	return []*survey.Question{
		{
			Name:   "Name",
			Prompt: &survey.Input{Message: "App Name?"},
		},
		{
			Name: "Version",
			Prompt: &survey.Select{
				Message: "Go Version",
				Options: cfg.Versions,
			},
		},
		{
			Name: "modules",
			Prompt: &survey.MultiSelect{
				Message: "choose module you need",
				Options: cfg.ModuleNames(),
			},
		},
	}
}

func ModuleQuestions(selectedModules []string, cfg config.Config) []*survey.Question {
	moduleQuestions := make([]*survey.Question, 0)
	for _, module := range selectedModules {
		moduleQuestions = append(moduleQuestions, &survey.Question{
			Name: module,
			Prompt: &survey.Select{
				Message: "choose technology for " + module,
				Options: cfg.ModuleOptions(module),
			},
		})
	}

	return moduleQuestions
}

func DockerfileQuestion() *survey.Confirm {
	return &survey.Confirm{
		Default: false,
		Message: "Do you want to create a Dockerfile?",
	}
}

func StartPointQuestion() *survey.Confirm {
	return &survey.Confirm{
		Message: "Do you want a cli app?",
		Default: false,
		Help:    "cli:yes, simple:no",
	}
}

func CliNameQuestion() *survey.Input {
	return &survey.Input{
		Message: "Name of CLI command",
		Default: "start",
	}
}

func handleErr(message string, err error) {
	if err != nil {
		panic(fmt.Errorf("%s: %w", message, err))
	}
}
