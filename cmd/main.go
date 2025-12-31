package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:     "sunju",
	Aliases: []string{"sunju-be"},
	Short:   "sunju-be",
	Long:    "Aplikasi registrasi dan tracking asset",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	rootCmd.Execute()
}

func init() {
	cobra.OnInitialize()
}
