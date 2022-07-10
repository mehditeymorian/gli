package config

func Default() Config {
	return Config{
		Versions: []string{
			"1.18",
			"1.17",
			"1.16",
			"1.15",
		},
		DB: []string{
			"none",
			"mysql",
			"postgres",
			"mongodb",
		},
		HTTP: []string{
			"none",
			"gofiber",
			"echo",
			"gin",
		},
		Log: []string{
			"none",
			"zap",
			"logrus",
		},
	}
}
