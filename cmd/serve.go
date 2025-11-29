package cmd

import (
	"galaveg/internal/bootstrap/http"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the http server of the application.",
	Long:  `Start the http server of the application.`,
	Run: func(cmd *cobra.Command, args []string) {
		http.Run()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
