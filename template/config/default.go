package config

import (
	{{if .HasHTTP}}"{{.Name}}/internal/http"{{end}}
)

func Default() Config {
	return Config{
		{{if .HasHTTP}}HTTP: http.Config{
			Port: "8080"
	},{{end}}
	}
}
