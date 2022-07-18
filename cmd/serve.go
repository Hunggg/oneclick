package cmd

import (
	"oneclick/api/http"
	"oneclick/config"

	"github.com/spf13/cobra"
)
var runHttpCmd = &cobra.Command{
	Use: "serve",
	Short: "Start HTTP server",
	Long: `Start HTTP server`,
	RunE: func (cmd *cobra.Command, args []string) error {
		var env config.Env
		env.LoadConfig()
		http.Init()
		return nil
	},
}