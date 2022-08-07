package config

import (
	"github.com/mehditeymorian/gli/internal/model"
)

func Default() Config {
	return Config{
		Versions: []string{
			"1.18",
			"1.17",
			"1.16",
			"1.15",
		},
		Modules: map[string][]model.Module{
			model.DB:         DBModules(),
			model.HTTP:       HTTPModules(),
			model.Logger:     LoggerModules(),
			model.Dockerfile: DockerfileModule(),
		},
		RequiredModules: []model.Module{
			{
				Name:        "config",
				DownloadURL: "template/config/",
				SavePath:    []string{"internal", "config"},
				Package:     []string{"github.com/knadh/koanf@latest", "github.com/tidwall/pretty@latest"},
				Files: []model.ModuleFile{
					{
						Name:           "config.go",
						RequireParsing: true,
					},
					{
						Name:           "default.go",
						RequireParsing: true,
					},
				},
			},
			{
				Name:        "Others",
				DownloadURL: "template/",
				SavePath:    nil,
				Package:     nil,
				Files: []model.ModuleFile{
					{
						Name:           "go.mod",
						RequireParsing: true,
					},
				},
			},
		},
	}
}

func DockerfileModule() []model.Module {
	return []model.Module{
		{
			Name:        "Dockerfile",
			DownloadURL: "template/",
			SavePath:    nil,
			Package:     nil,
			Files: []model.ModuleFile{
				{
					Name:           "Dockerfile",
					RequireParsing: false,
				},
			},
		},
	}
}

func LoggerModules() []model.Module {
	return []model.Module{
		{
			Name:  model.None,
			Files: nil,
		},
		{
			Name:        "zap",
			DownloadURL: "template/logger/zap/",
			SavePath:    []string{"internal", "logger"},
			Package:     []string{"go.uber.org/zap@latest"},
			Files: []model.ModuleFile{
				{
					Name:           "config.go",
					RequireParsing: false,
				},
				{
					Name:           "log.go",
					RequireParsing: false,
				},
			},
		},
	}
}

func HTTPModules() []model.Module {
	return []model.Module{
		{
			Name:  model.None,
			Files: nil,
		},
		{
			Name:  "gofiber",
			Files: nil,
		},
		{
			Name:  "echo",
			Files: nil,
		},
		{
			Name:  "gin",
			Files: nil,
		},
	}
}

func DBModules() []model.Module {
	return []model.Module{
		{
			Name:  model.None,
			Files: nil,
		},
		{
			Name:  "mysql",
			Files: nil,
		},
		{
			Name:  "postgres",
			Files: nil,
		},
		{
			Name:        "mongo",
			DownloadURL: "template/db/mongo/",
			SavePath:    []string{"internal", "db"},
			Package:     []string{"go.mongodb.org/mongo-driver/mongo"},
			Files: []model.ModuleFile{
				{
					Name:           "config.go",
					RequireParsing: false,
				},
				{
					Name:           "mongo.go",
					RequireParsing: false,
				},
				{
					Name:           "mongo_test.go",
					RequireParsing: false,
				},
			},
		},
	}
}
