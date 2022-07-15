package model

import "github.com/mehditeymorian/gli/internal/config"

type App struct {
	Modules []string

	Name       string
	Version    string
	Logger     string
	DB         string
	HTTP       string
	Dockerfile bool
}

func EmptyApp() *App {
	return &App{
		Logger:     config.None,
		DB:         config.None,
		HTTP:       config.None,
		Dockerfile: false,
	}
}
