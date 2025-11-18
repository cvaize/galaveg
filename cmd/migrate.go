package cmd

import (
	"database/sql"
	"fmt"
	"galaveg/config"
	"galaveg/connections/db"
	"galaveg/database/migrations"
	"galaveg/utils/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"path/filepath"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Start database migrations.",
	Long:  `Start database migrations.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		C := config.New(filepath.Join(viper.GetString("APP_FOLDER"), ".env"))
		DB := db.New(C.Db)
		err := UpMigration(C, DB)
		logger.Infof("The migration was successful!")
		return err
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func createMigrationsTable(c *config.Config, db *sql.DB) error {
	//goland:noinspection ALL
	query := `CREATE TABLE IF NOT EXISTS ` + c.Db.Prefix + `_migrations (
		id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(255) NOT NULL UNIQUE
	);`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

type migrationRow struct {
	id   uint64
	name string
}

func LoadMigrations(c *config.Config, db *sql.DB) ([]migrationRow, error) {
	//goland:noinspection ALL
	query := "SELECT * FROM " + c.Db.Prefix + "_migrations ORDER BY id ASC;"
	rows, err := db.Query(query)
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

func InsertMigration(c *config.Config, db *sql.DB, name string) error {
	//goland:noinspection ALL
	query := "INSERT INTO " + c.Db.Prefix + "_migrations (name) VALUES (?);"
	_, err := db.Exec(query, name)
	if err != nil {
		return err
	}

	return nil
}

func UpMigration(c *config.Config, db *sql.DB) error {
	err := createMigrationsTable(c, db)
	if err != nil {
		return err
	}

	migrationRows, err := LoadMigrations(c, db)
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
			err1 := m.Up(c, db)
			if err1 != nil {
				return err1
			}
			err2 := InsertMigration(c, db, m.Uuid)
			if err2 != nil {
				return err2
			}
			logger.Infof(fmt.Sprintf("Migrated - %s", m.Uuid))
		}
	}

	return nil
}
