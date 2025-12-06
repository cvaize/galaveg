package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// keyGenerateCmd represents the keyGenerate command
var keyGenerateCmd = &cobra.Command{
	Use:   "key:generate",
	Short: "Set the application key",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("keyGenerate called")
	},
}

func init() {
	rootCmd.AddCommand(keyGenerateCmd)

	//keyGenerateCmd.Flags().BoolP("show", "s", false, "Display the key instead of modifying files")
	//keyGenerateCmd.Flags().BoolP("force", "f", false, "Force the operation to run when in production")
}

func generateRandomKey() {

}
