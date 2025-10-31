package cmd

import (
	"errors"
	"fmt"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"os"
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
		camelName := lo.CamelCase(originalName)
		fmt.Println(snakeName)
		fmt.Println(camelName)

		wd, err := os.Getwd()
		if err != nil {
			return errors.New(err.Error())
		}

		migrationFileName := fmt.Sprintf("%s/database/migrations/%s_%s.go", wd, time.Now().Format("2006_01_02_150405"), snakeName)

		if _, err := os.Stat(migrationFileName); err == nil {
			return errors.New(fmt.Sprintf("the %s file already exists", migrationFileName))
		}

		file, err := os.Create(migrationFileName)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		migrationContent := getMakeMigrationTemplate()
		_, err = file.WriteString(migrationContent)
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
`
}
