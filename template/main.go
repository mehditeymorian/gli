package main

import (
	"{{ .Name }}/internal/config"
	{{if .HasLogger}}"{{ .Name }}/internal/logger"{{end}}
	{{if .HasHTTP}}"{{ .Name }}/internal/http/app"{{end}}
)

func main() {
	cfg := config.Load("config.yaml")
	{{if .HasLogger}}log := logger.New(cfg.Logger){{end}}
	{{if .HasHTTP}}app := app.New(cfg.HTTP){{end}}

	{{if .HasHTTP}}app.Serve(){{end}}
}
