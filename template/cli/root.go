package cmd

import (
	"github.com/spf13/cobra"
	"{{ .Name }}/internal/cmd"
)

func Execute() {

	root := &cobra.Command{
		Use: "{{ .ShortName}}",
		Short: ``,
		Long: ``,
	}

	root.AddCommands(
		cmd.{{ .CliName | capitalize }},
		)


	if err := root.Execute(); err != nil {
		panic(err)
	}
}
