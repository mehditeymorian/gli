package config

import (
	{{if .HasHTTP}}"{{.Name}}/internal/http"{{end}}
	{{if .HasLogger}}"{{.Name}}/internal/logger"{{end}}
)

func Default() Config {
	return Config{
		{{if .HasHTTP}}HTTP: http.Config{
			Port: "8080",
	},{{end}}
		{{if .HasLogger}}Logger: logger.Config{
			Level: "warn",
	},{{end}}
		{{if .HasDB}}DB: db.Config{
			URI: "mongodb:localhost:27017",
			Name: "{{.ShortName }}",
	},{{end}}
	}
}
