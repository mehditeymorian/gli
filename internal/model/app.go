package model

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
	return &App{}
}
