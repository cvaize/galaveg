package cmd

import (
	"fmt"
	"galaveg/connections"
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
		err := upMigration()
		logger.Infof("The migration was successful!")
		return err
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func createMigrationsTable() error {
	//goland:noinspection ALL
	query := `CREATE TABLE IF NOT EXISTS __migrations (
		id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(255) NOT NULL UNIQUE
	);`
	_, err := connections.DB.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

type migrationRow struct {
	id   uint64
	name string
}

func loadMigrations() ([]migrationRow, error) {
	//goland:noinspection ALL
	query := "SELECT * FROM __migrations ORDER BY id ASC;"
	rows, err := connections.DB.Query(query)
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

func insertMigration(name string) error {
	//goland:noinspection ALL
	query := `INSERT INTO __migrations (name) VALUES (?);`
	_, err := connections.DB.Exec(query, name)
	if err != nil {
		return err
	}

	return nil
}

func upMigration() error {
	err := createMigrationsTable()
	if err != nil {
		return err
	}

	migrationRows, err := loadMigrations()
	if err != nil {
		return err
	}

	index := make(map[string]uint64)

	for _, m := range migrationRows {
		index[m.name] = m.id
	}

	for _, m := range migrations.Migrations {
		if _, ok := index[m.Uuid]; !ok {
			logger.Infof(fmt.Sprintf("Migrating - %s", m.Uuid))
			err1 := m.Up()
			if err1 != nil {
				return err1
			}
			err2 := insertMigration(m.Uuid)
			if err2 != nil {
				return err2
			}
			logger.Infof(fmt.Sprintf("Migrated - %s", m.Uuid))
		}
	}

	return nil
}
