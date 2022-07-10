package config

type Config struct {
	Versions []string
	Modules  map[string][]string
}

func (c Config) ModuleNames() []string {
	modules := make([]string, 0)
	for module, _ := range c.Modules {
		modules = append(modules, module)
	}

	return modules
}

func Load() Config {
	return Default()
}
