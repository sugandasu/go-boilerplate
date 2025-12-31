package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sugandasu/go-boilerplate/config"
	"github.com/sugandasu/go-boilerplate/internal/app/server"
)

func initRestCmd() *cobra.Command {
	return &cobra.Command{
		Use:  "rest",
		Long: "REST API",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := config.Load()
			server.RunRestServer(cfg)
		},
	}
}

func init() {
	server := &cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	server.AddCommand(initRestCmd())
	rootCmd.AddCommand(server)
}
