package model

import (
	"strings"
)

type App struct {
	Modules []string

	Name       string
	ShortName  string
	Version    string
	Logger     string
	DB         string
	HTTP       string
	Dockerfile bool
}

func EmptyApp() *App {
	return &App{
		Logger:     None,
		DB:         None,
		HTTP:       None,
		Dockerfile: false,
	}
}

func (a *App) Params() map[string]any {
	return map[string]any{
		"Name":      a.Name,
		"ShortName": a.ShortName,
		"Version":   a.Version,
		"HasLogger": a.Logger != None,
		"HasDB":     a.DB != None,
		"HasHTTP":   a.HTTP != None,
	}
}

// Execute extract some fields from existing fields
func (a *App) Execute() {
	if strings.Contains(a.Name, "/") {
		split := strings.Split(a.Name, "/")
		a.ShortName = split[len(split)-1]
	} else {
		a.ShortName = a.Name
	}
}
