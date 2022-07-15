package config

import "github.com/mehditeymorian/gli/internal/model"

type Config struct {
	Versions []string
	Modules  map[string][]model.Module
}

func (c Config) ModuleNames() []string {
	modules := make([]string, 0)
	for module, _ := range c.Modules {
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

func Load() Config {
	return Default()
}
