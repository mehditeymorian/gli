package main

import (
	"{{ .Name }}/internal/config"
	{{if .HasLogger}}"{{ .Name }}/internal/logger"{{end}}
)

func main() {
	cfg := config.Load("config.yaml")
	{{if .HasLogger}}log := logger.New(cfg.Logger){{end}}
}
