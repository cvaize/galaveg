package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// migrateRollbackCmd represents the migrate command
var migrateRollbackCmd = &cobra.Command{
	Use:   "migrate:rollback",
	Short: "Rollback the database migration.",
	Long:  `Rollback the database migration.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migrate:rollback")
	},
}

func init() {
	rootCmd.AddCommand(migrateRollbackCmd)
}
