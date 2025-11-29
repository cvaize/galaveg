package cmd

import (
	"galaveg/internal/bootstrap/http"

	"github.com/spf13/cobra"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Start the http server of the application.",
	Long:  `Start the http server of the application.`,
	Run: func(cmd *cobra.Command, args []string) {
		http.Run()
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
