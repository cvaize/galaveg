package cmd

import (
	"galaveg/internal/bootstrap/serve"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the http server and chat of the application.",
	Long:  `Start the http server and chat of the application.`,
	Run: func(cmd *cobra.Command, args []string) {
		serve.Run()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
