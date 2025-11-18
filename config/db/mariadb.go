package db

import (
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("MARIADB_DRIVER", "mysql")
	viper.SetDefault("MARIADB_HOST", "localhost")
	viper.SetDefault("MARIADB_PORT", 3306)
	viper.SetDefault("MARIADB_DATABASE", "database")
	viper.SetDefault("MARIADB_USERNAME", "database_user")
	viper.SetDefault("MARIADB_PASSWORD", "database_password")
	viper.SetDefault("MARIADB_PREFIX", "")
}

type MariadbConfig struct {
	Driver   string
	Host     string
	Port     int
	Database string
	Username string
	Password string
	Prefix   string
}

func (c *MariadbConfig) DSN() string {
	// TODO: Подумать над параметрами, на вроде multiStatements
	//"user:password@tcp(localhost:3306)/database?parseTime=true"
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		c.Username, c.Password, c.Host, c.Port, c.Database,
	)
}

func NewMariadbConfig() MariadbConfig {
	return MariadbConfig{
		Driver:   viper.GetString("MARIADB_DRIVER"),
		Host:     viper.GetString("MARIADB_HOST"),
		Port:     viper.GetInt("MARIADB_PORT"),
		Database: viper.GetString("MARIADB_DATABASE"),
		Username: viper.GetString("MARIADB_USERNAME"),
		Password: viper.GetString("MARIADB_PASSWORD"),
		Prefix:   viper.GetString("MARIADB_PREFIX"),
	}
}
