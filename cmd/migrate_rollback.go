package cmd

import (
	"database/sql"
	"fmt"
	"galaveg/config"
	"galaveg/connections/db"
	"galaveg/database/migrations"
	"galaveg/utils/logger"
	"github.com/samber/lo"
	"github.com/spf13/viper"
	"path/filepath"

	"github.com/spf13/cobra"
)

// migrateRollbackCmd represents the migrate command
var migrateRollbackCmd = &cobra.Command{
	Use:   "migrate:rollback",
	Short: "Rollback the database migration.",
	Long:  `Rollback the database migration.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		C := config.New(filepath.Join(viper.GetString("APP_FOLDER"), ".env"))
		DB := db.New(C.Db)
		err := RollbackMigration(C, DB)
		logger.Infof("The rollback migration was successful!")
		return err
	},
}

func init() {
	rootCmd.AddCommand(migrateRollbackCmd)
}

func DeleteMigration(c *config.Config, db *sql.DB, name string) error {
	//goland:noinspection ALL
	query := "DELETE FROM " + c.Db.Prefix + "_migrations WHERE name=?"
	_, err := db.Exec(query, name)
	if err != nil {
		return err
	}

	return nil
}

func RollbackMigration(c *config.Config, db *sql.DB) error {
	err := createMigrationsTable(c, db)
	if err != nil {
		return err
	}

	migrationRows, err := LoadMigrations(c, db)
	if err != nil {
		return err
	}

	lastMigrationRow, exists := lo.Last(migrationRows)

	if !exists {
		return nil
	}

	ms := migrations.GetMigrations()

	for _, m := range ms {
		if lastMigrationRow.name == m.Uuid {
			logger.Infof(fmt.Sprintf("Rollback migrating - %s", m.Uuid))
			err1 := m.Down(c, db)
			if err1 != nil {
				return err1
			}
			err2 := DeleteMigration(c, db, m.Uuid)
			if err2 != nil {
				return err2
			}
			logger.Infof(fmt.Sprintf("Rollback migrated - %s", m.Uuid))
		}
	}

	return nil
}
