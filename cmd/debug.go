package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// debugCmd represents the debug command
var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "debug",
	Long:  `debug`,
	Run: func(cmd *cobra.Command, args []string) {
		//https://github.com/gin-contrib/i18n
		//https://github.com/nicksnyder/go-i18n
		//https://www.squash.io/implementing-internationalization-in-gin-with-go-libraries/
		fmt.Println("debug called")
	},
}

func init() {
	rootCmd.AddCommand(debugCmd)
}
