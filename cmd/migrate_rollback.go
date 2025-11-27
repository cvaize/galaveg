package cmd

import (
	"fmt"
	"galaveg/bootstrap/providers"
	"galaveg/config"
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
		cfg := config.MustDefault()
		ctx := providers.MustContext(cfg)
		err := RollbackMigration(ctx)
		logger.Infof("The rollback migration was successful!")
		return err
	},
}

func init() {
	rootCmd.AddCommand(migrateRollbackCmd)
}

func DeleteMigration(ctx *providers.Context, name string) error {
	//goland:noinspection ALL
	query := "DELETE FROM " + ctx.Cfg.Db.Prefix + "_migrations WHERE name=?"
	_, err := ctx.Infra.Db.Exec(query, name)
	if err != nil {
		return err
	}

	return nil
}

func RollbackMigration(ctx *providers.Context) error {
	err := createMigrationsTable(ctx)
	if err != nil {
		return err
	}

	migrationRows, err := LoadMigrations(ctx)
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
			err1 := m.Down(ctx)
			if err1 != nil {
				return err1
			}
			err2 := DeleteMigration(ctx, m.Uuid)
			if err2 != nil {
				return err2
			}
			logger.Infof(fmt.Sprintf("Rollback migrated - %s", m.Uuid))
		}
	}

	return nil
}
