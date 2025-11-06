package cmd

import (
	"errors"
	"fmt"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"time"
)

// makeMigrationCmd represents the makeMigration command
var makeMigrationCmd = &cobra.Command{
	Use:                   "make:migration [name]",
	Short:                 "Create a migration.",
	Long:                  `Create a migration.`,
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		originalName, err := lo.Nth(args, 0)
		if err != nil {
			return errors.New("the [name] argument is not specified")
		}
		snakeName := lo.SnakeCase(originalName)
		pascalName := lo.PascalCase(originalName)

		wd, err := os.Getwd()
		if err != nil {
			return errors.New(err.Error())
		}

		now := time.Now()
		migrationName := fmt.Sprintf("%s_%s", now.Format("2006_01_02_150405"), snakeName)
		migrationFileName := fmt.Sprintf("%s/database/migrations/%s.go", wd, migrationName)

		funcName := pascalName + now.Format("20060102150405")
		// CreateUsersTable00010101000000Up
		upFuncName := funcName + "Up"

		// CreateUsersTable00010101000000Down
		downFuncName := funcName + "Down"

		if _, err := os.Stat(migrationFileName); err == nil {
			return errors.New(fmt.Sprintf("the %s file already exists", migrationFileName))
		}

		file, err := os.Create(migrationFileName)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		migrationContent := strings.ReplaceAll(getMakeMigrationTemplate(), "{{ upName }}", upFuncName)
		migrationContent = strings.ReplaceAll(migrationContent, "{{ downName }}", downFuncName)
		_, err = file.WriteString(migrationContent)
		if err != nil {
			return errors.New(err.Error())
		}

		migrationsFileName := fmt.Sprintf("%s/database/migrations/migrations.go", wd)
		migrationsContentBytes, err := os.ReadFile(migrationsFileName)
		if err != nil {
			panic(err)
		}

		migrationsContent := string(migrationsContentBytes)

		newMigration := `Migration{
			Uuid: "` + migrationName + `",
			Up:   ` + upFuncName + `,
			Down: ` + downFuncName + `,
		},
		// NEW_MIGRATIONS_TAG`

		updatedMigrationsContent := strings.ReplaceAll(migrationsContent, "// NEW_MIGRATIONS_TAG", newMigration)

		migrationsFile, err := os.Create(migrationsFileName)
		if err != nil {
			panic(err)
		}
		defer migrationsFile.Close()

		_, err = migrationsFile.WriteString(updatedMigrationsContent)
		if err != nil {
			return errors.New(err.Error())
		}

		fmt.Printf("%s created at %s\n", originalName, migrationFileName)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(makeMigrationCmd)
}

func getMakeMigrationTemplate() string {
	return `package migrations

import (
	"database/sql"
	"galaveg/config"
)

func {{ upName }}(c *config.Config, db *sql.DB) error {
	return nil
}

func {{ downName }}(c *config.Config, db *sql.DB) error {
	return nil
}
`
}
