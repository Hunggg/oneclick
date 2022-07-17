package cmd

import (
	"oneclick/config"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "serve",
	Short: "start http server with configured api",
	Long:  `Starts a http server and serves the configured api`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var env config.Env
		env.LoadConfig()
		return nil
	},
}