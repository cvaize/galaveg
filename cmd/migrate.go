package cmd

import (
	"fmt"
	"galaveg/bootstrap/providers"
	"galaveg/config"
	"galaveg/database/migrations"
	"galaveg/utils/logger"
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Start database migrations.",
	Long:  `Start database migrations.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		C := config.MustDefaultConfig()
		ctx := providers.MustContext(C)
		err := UpMigration(ctx)
		logger.Infof("The migration was successful!")
		return err
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func createMigrationsTable(ctx *providers.Context) error {
	//goland:noinspection ALL
	query := `CREATE TABLE IF NOT EXISTS ` + ctx.C.Db.Prefix + `_migrations (
		id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(255) NOT NULL UNIQUE
	);`
	_, err := ctx.DB.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

type migrationRow struct {
	id   uint64
	name string
}

func LoadMigrations(ctx *providers.Context) ([]migrationRow, error) {
	//goland:noinspection ALL
	query := "SELECT * FROM " + ctx.C.Db.Prefix + "_migrations ORDER BY id ASC;"
	rows, err := ctx.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	migrationRows := make([]migrationRow, 0)

	for rows.Next() {
		m := migrationRow{}
		err := rows.Scan(&m.id, &m.name)
		if err != nil {
			return nil, err
		}
		migrationRows = append(migrationRows, m)
	}

	return migrationRows, err
}

func InsertMigration(ctx *providers.Context, name string) error {
	//goland:noinspection ALL
	query := "INSERT INTO " + ctx.C.Db.Prefix + "_migrations (name) VALUES (?);"
	_, err := ctx.DB.Exec(query, name)
	if err != nil {
		return err
	}

	return nil
}

func UpMigration(ctx *providers.Context) error {
	err := createMigrationsTable(ctx)
	if err != nil {
		return err
	}

	migrationRows, err := LoadMigrations(ctx)
	if err != nil {
		return err
	}

	index := make(map[string]uint64)

	for _, m := range migrationRows {
		index[m.name] = m.id
	}

	ms := migrations.GetMigrations()

	for _, m := range ms {
		if _, ok := index[m.Uuid]; !ok {
			logger.Infof(fmt.Sprintf("Migrating - %s", m.Uuid))
			err1 := m.Up(ctx)
			if err1 != nil {
				return err1
			}
			err2 := InsertMigration(ctx, m.Uuid)
			if err2 != nil {
				return err2
			}
			logger.Infof(fmt.Sprintf("Migrated - %s", m.Uuid))
		}
	}

	return nil
}
