package cmd

import (
	"github.com/spf13/cobra"
	"{{ .Name }}/internal/config"
	"{{ .Name }}/internal/logger"
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
	cfg := config.Load()
	{{if .HasLogger}}log := logger.New(cfg.Logger){{end}}
}
