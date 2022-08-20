package cmd

import (
	"github.com/spf13/cobra"
	"{{ .Name }}/internal/config"
	{{if .HasLogger}}"{{ .Name }}/internal/logger"{{end}}
	{{if .HasHTTP}}"{{ .Name }}/internal/http/app"{{end}}
	{{if .HasDB}}"{{ .Name }}/internal/db"{{end}}
)

func {{ capitalize .CliName }}() *cobra.Command {

	command := &cobra.Command{
		Use: "{{ .CliName }}",
		Short : ``,
		Long: ``,
		Run: main,
	}

	return command
}

func main(cmd *cobra.Command, args []string)  {
	cfg := config.Load("config.yaml")

	{{if .HasLogger}}log := logger.New(cfg.Logger){{end}}

	{{if .HasDB}}
	{{if eq .DB "mongo"}} database, err := db.Connect(cfg.DB){{end}}

	if err != nil{
		log.Fatalf("failed to connect to db: %v", err)
	}
	{{end}}

	{{if .HasHTTP}}app := app.New(cfg.HTTP){{end}}

	{{if .HasHTTP}}app.Serve(){{end}}
}
