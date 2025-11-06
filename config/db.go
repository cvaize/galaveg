package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("DB_DRIVER", "mysql")
	viper.SetDefault("DB_HOST", "mysql")
	viper.SetDefault("DB_PORT", "3306")
	viper.SetDefault("DB_DATABASE", "database")
	viper.SetDefault("DB_USERNAME", "database_user")
	viper.SetDefault("DB_PASSWORD", "database_password")
	viper.SetDefault("DB_PREFIX", "")
}

type DbConfig struct {
	Driver   string
	Host     string
	Port     string
	Database string
	Username string
	Password string
	Prefix   string
}

func (c DbConfig) DSN() string {
	// TODO: Подумать над параметрами, на вроде multiStatements
	//"user:password@tcp(localhost:3306)/database?parseTime=true"
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		c.Username, c.Password, c.Host, c.Port, c.Database,
	)
}

func NewDbConfig() DbConfig {
	return DbConfig{
		Driver:   viper.GetString("DB_DRIVER"),
		Host:     viper.GetString("DB_HOST"),
		Port:     viper.GetString("DB_PORT"),
		Database: viper.GetString("DB_DATABASE"),
		Username: viper.GetString("DB_USERNAME"),
		Password: viper.GetString("DB_PASSWORD"),
		Prefix:   viper.GetString("DB_PREFIX"),
	}
}
