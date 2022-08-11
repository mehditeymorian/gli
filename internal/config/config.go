package config

import (
	"github.com/mehditeymorian/gli/internal/model"
)

type Config struct {
	Versions          []string
	SelectableModules map[string]model.ModuleGroup
	RequiredModules   []model.Module
}

func (c Config) ModuleNames() []string {
	modules := make([]string, 0)
	for module, val := range c.SelectableModules {

		if !val.Selectable {
			continue
		}

		modules = append(modules, module)
	}

	return modules
}

func (c Config) ModuleOptions(name string) []string {
	names := make([]string, 0)

	for _, module := range c.SelectableModules[name].Modules {
		names = append(names, module.Name)
	}

	return names
}

func (c Config) GetRequiredModules(app *model.SurveyResult) []model.Module {
	modules := make([]model.Module, 0)

	if app.HTTP != model.None {
		module := c.SearchModule(model.HTTP, app.HTTP)
		if module != nil {
			modules = append(modules, *module)
		}
	}

	if app.DB != model.None {
		module := c.SearchModule(model.DB, app.DB)
		if module != nil {
			modules = append(modules, *module)
		}
	}

	if app.Logger != model.None {
		module := c.SearchModule(model.Logger, app.Logger)
		if module != nil {
			modules = append(modules, *module)
		}
	}

	if app.CliType {
		module := c.SearchModule(model.StartPoint, model.StartPointCli)
		if module != nil {
			modules = append(modules, *module)
		}
	} else {
		module := c.SearchModule(model.StartPoint, model.StartPointSimple)
		if module != nil {
			modules = append(modules, *module)
		}
	}

	if app.Dockerfile {
		module := c.SearchModule(model.Dockerfile, model.Dockerfile)
		modules = append(modules, *module)
	}

	return modules
}

func (c Config) SearchModule(module, technology string) *model.Module {
	for _, m := range c.SelectableModules[module].Modules {
		if m.Name == technology {
			return &m
		}
	}

	return nil
}

func Load() Config {
	return Default()
}
