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
		Modules: map[string]model.ModuleGroup{
			model.DB:         DBModules(),
			model.HTTP:       HTTPModules(),
			model.Logger:     LoggerModules(),
			model.Dockerfile: DockerfileModule(),
		},
		RequiredModules: []model.Module{
			configModule(),
			modModule(),
		},
	}
}

func modModule() model.Module {
	return model.Module{
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
	}
}

func configModule() model.Module {
	return model.Module{
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
	}
}

func DockerfileModule() model.ModuleGroup {
	modules := []model.Module{
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
	return model.ModuleGroup{
		Selectable: false,
		Modules:    modules,
	}
}

func LoggerModules() model.ModuleGroup {
	modules := []model.Module{
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
	return model.ModuleGroup{
		Selectable: true,
		Modules:    modules,
	}
}

func HTTPModules() model.ModuleGroup {
	modules := []model.Module{
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
	return model.ModuleGroup{
		Selectable: true,
		Modules:    modules,
	}
}

func DBModules() model.ModuleGroup {
	modules := []model.Module{
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

	return model.ModuleGroup{
		Selectable: true,
		Modules:    modules,
	}
}
