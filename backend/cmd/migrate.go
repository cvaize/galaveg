package cmd

import (
	"galaveg/config"
	"galaveg/connections"
	"github.com/spf13/cobra"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Запуск миграции базы данных.",
	Long:  `Запуск миграции базы данных.`,
	Run: func(cmd *cobra.Command, args []string) {
		driver, _ := mysql.WithInstance(connections.MySQL, &mysql.Config{})
		m, _ := migrate.NewWithDatabaseInstance(
			"file:///migrations/mysql",
			config.Config.Mysql.Database,
			driver,
		)

		//m.Steps(1)
		if err := m.Up(); err != nil {
			panic(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

//
///**
// * The name and signature of the console command.
// *
// * @var string
// */
//protected $signature = 'migrate {--database= : The database connection to use}
//{--force : Force the operation to run when in production}
//{--path=* : The path(s) to the migrations files to be executed}
//{--realpath : Indicate any provided migration file paths are pre-resolved absolute paths}
//{--schema-path= : The path to a schema dump file}
//{--pretend : Dump the SQL queries that would be run}
//{--seed : Indicates if the seed task should be re-run}
//{--seeder= : The class name of the root seeder}
//{--step : Force the migrations to be run so they can be rolled back individually}
//{--graceful : Return a successful exit code even if an error occurs}';
//
///**
// * The console command description.
// *
// * @var string
// */
//protected $description = 'Run the database migrations';
