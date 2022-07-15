package config

import "github.com/mehditeymorian/gli/internal/model"

const (
	None   = "none"
	DB     = "db"
	HTTP   = "http"
	Logger = "logger"
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
			DB: {
				{
					Name:  None,
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
			HTTP: {
				{
					Name:  None,
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
			Logger: {
				{
					Name:  None,
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
