package cmd

import "github.com/spf13/cobra"

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

}
