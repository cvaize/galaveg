package cmd

import (
	"errors"
	"fmt"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"os"
	"regexp"
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
		defer file.Close()
		if err != nil {
			panic(err)
		}

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

		newMigration := `
	Migration{
		Uuid: "` + migrationName + `",
		Up:   ` + upFuncName + `,
		Down: ` + downFuncName + `,
	},`

		substring := "var Migrations = []Migration{}"
		updatedMigrationsContent := migrationsContent
		if strings.Contains(migrationsContent, substring) {
			template := `var Migrations = []Migration{` + newMigration + `
}`
			updatedMigrationsContent = strings.ReplaceAll(migrationsContent, substring, template)
		} else {
			re := regexp.MustCompile(`(?s)(var\s+Migrations\s*=\s*\[\]\s*Migration\s*{)(.*?)(\n}\s*)`)
			updatedMigrationsContent = re.ReplaceAllString(migrationsContent, "${1}${2}"+newMigration+"${3}")
		}

		migrationsFile, err := os.Create(migrationsFileName)
		defer migrationsFile.Close()
		if err != nil {
			panic(err)
		}

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
)

func {{ upName }}(db *sql.DB) error {
	return nil
}

func {{ downName }}(db *sql.DB) error {
	return nil
}
`
}
