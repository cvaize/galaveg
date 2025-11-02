package cmd

import (
	"fmt"
	"galaveg/connections"
	"galaveg/database/migrations"
	"galaveg/utils/logger"
	"github.com/samber/lo"

	"github.com/spf13/cobra"
)

// migrateRollbackCmd represents the migrate command
var migrateRollbackCmd = &cobra.Command{
	Use:   "migrate:rollback",
	Short: "Rollback the database migration.",
	Long:  `Rollback the database migration.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := rollbackMigration()
		logger.Infof("The rollback migration was successful!")
		return err
	},
}

func init() {
	rootCmd.AddCommand(migrateRollbackCmd)
}

func deleteMigration(name string) error {
	//goland:noinspection ALL
	query := `DELETE FROM __migrations WHERE name=?`
	_, err := connections.DB.Exec(query, name)
	if err != nil {
		return err
	}

	return nil
}

func rollbackMigration() error {
	err := createMigrationsTable()
	if err != nil {
		return err
	}

	migrationRows, err := loadMigrations()
	if err != nil {
		return err
	}

	lastMigrationRow, exists := lo.Last(migrationRows)

	if !exists {
		return nil
	}

	for _, m := range migrations.Migrations {
		if lastMigrationRow.name == m.Uuid {
			logger.Infof(fmt.Sprintf("Rollback migrating - %s", m.Uuid))
			err1 := m.Down()
			if err1 != nil {
				return err1
			}
			err2 := deleteMigration(m.Uuid)
			if err2 != nil {
				return err2
			}
			logger.Infof(fmt.Sprintf("Rollback migrated - %s", m.Uuid))
		}
	}

	return nil
}
