package config

import (
	"github.com/mehditeymorian/gli/internal/model"
)

type Config struct {
	Versions        []string
	Modules         map[string][]model.Module
	RequiredModules []model.Module
}

func (c Config) ModuleNames() []string {
	modules := make([]string, 0)
	for module, val := range c.Modules {
		// modules with 1 option are yes or no type of questions, and they are mandatory to ask from user.
		if len(val) < 2 {
			continue
		}

		modules = append(modules, module)
	}

	return modules
}

func (c Config) ModuleOptions(name string) []string {
	names := make([]string, 0)

	for _, module := range c.Modules[name] {
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

	if app.Dockerfile {
		module := c.SearchModule(model.Dockerfile, model.Dockerfile)
		modules = append(modules, *module)
	}

	return modules
}

func (c Config) SearchModule(module, technology string) *model.Module {
	for _, m := range c.Modules[module] {
		if m.Name == technology {
			return &m
		}
	}

	return nil
}

func Load() Config {
	return Default()
}
