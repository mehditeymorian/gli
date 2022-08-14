package main

import (
	"{{ .Name }}/internal/config"
	"{{ .Name }}/internal/logger"
)

func main() {
	cfg := config.Load("config.yaml")
	{{if .HasLogger}}log := logger.New(cfg.Logger){{end}}
}
