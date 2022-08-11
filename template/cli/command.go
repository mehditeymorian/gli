package cmd

import "github.com/spf13/cobra"

func {{ .CliName | capitalize }}() *cobra.Command {

	command := &cobra.Command{
		Use: "{{ .CliName }}",
		Short : ``,
		Long: ``,
		Run: main,
	}
}

func main(cmd *cobra.Command, args []string)  {

}
