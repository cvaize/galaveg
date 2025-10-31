package cmd

import (
	"galaveg/bootstrap"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the http server of the application.",
	Long:  `Start the http server of the application.`,
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.Http()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
