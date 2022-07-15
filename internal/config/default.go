package config

func Default() Config {
	return Config{
		Versions: []string{
			"1.18",
			"1.17",
			"1.16",
			"1.15",
		},
		Modules: map[string][]Module{
			"db": {
				{
					Name:  "none",
					Files: nil,
				},
				{
					Name:  "mysql",
					Files: []string{},
				},
				{
					Name:  "postgres",
					Files: []string{},
				},
				{
					Name:  "mongodb",
					Files: []string{},
				},
			},
			"http": {
				{
					Name:  "none",
					Files: []string{},
				},
				{
					Name:  "gofiber",
					Files: []string{},
				},
				{
					Name:  "echo",
					Files: []string{},
				},
				{
					Name:  "gin",
					Files: []string{},
				},
			},
			"logger": {
				{
					Name:  "none",
					Files: []string{},
				},
				{
					Name: "zap",
					Files: []string{
						"config.go",
						"log.go",
					},
				},
				{
					Name:  "logrus",
					Files: []string{},
				},
			},
		},
	}
}
