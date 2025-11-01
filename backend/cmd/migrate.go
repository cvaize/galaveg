package cmd

import (
	"database/sql"
	"fmt"
	"galaveg/connections"
	"galaveg/database/migrations"
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Start database migrations.",
	Long:  `Start database migrations.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return upMigration(connections.DB)
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func createMigrationsTable(db *sql.DB) error {
	//goland:noinspection SqlNoDataSourceInspection
	query := `CREATE TABLE IF NOT EXISTS __migrations (
		id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(255) NOT NULL UNIQUE
	);`
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return nil
}

func upMigration(db *sql.DB) error {
	err := createMigrationsTable(db)
	if err != nil {
		return err
	}

	fmt.Println(migrations.Migrations)

	return nil
}
