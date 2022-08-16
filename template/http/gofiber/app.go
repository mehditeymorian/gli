package app

import (
	"{{ .Name }}/internal/http"
)

type App struct {
	App  fiber.App
	Port string
}

func New(cfg http.Config) App {
	app := fiber.New()

	// register handlers here

	return App{
		App:  app,
		Port: cfg.Port,
	}
}

func (a App) Serve() {
	if err := a.App.Listen(":" + a.Port); err != nil {
		panic(err)
	}
}
