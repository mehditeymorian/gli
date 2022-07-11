package config

func Default() Config {
	return Config{
		Versions: []string{
			"1.18",
			"1.17",
			"1.16",
			"1.15",
		},
		Modules: map[string][]string{
			"db": {
				"none",
				"mysql",
				"postgres",
				"mongodb",
			},
			"http": {
				"none",
				"gofiber",
				"echo",
				"gin",
			},
			"logger": {
				"none",
				"zap",
				"logrus",
			},
		},
	}
}
