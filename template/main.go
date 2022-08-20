package main

import (
	"log"

	"{{ .Name }}/internal/config"
	{{if .HasLogger}}"{{ .Name }}/internal/logger"{{end}}
	{{if .HasHTTP}}"{{ .Name }}/internal/http/app"{{end}}
	{{if .HasDB}}"{{ .Name }}/internal/db"{{end}}
)

func main() {
	cfg := config.Load("config.yaml")

	{{if .HasLogger}}logger := logger.New(cfg.Logger){{end}}

	{{if .HasDB}}
	{{if eq .DB "mongo"}} database, err := db.Connect(cfg.DB){{end}}

	if err != nil{
		log.Fatalf("failed to connect to db: %v", err)
	}
	{{end}}

	{{if .HasHTTP}}app := app.New(cfg.HTTP){{end}}

	{{if .HasHTTP}}app.Serve(){{end}}
}
