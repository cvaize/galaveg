package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("MYSQL_HOST", "mysql")
	viper.SetDefault("MYSQL_PORT", "3306")
	viper.SetDefault("MYSQL_DATABASE", "database")
	viper.SetDefault("MYSQL_USERNAME", "database_user")
	viper.SetDefault("MYSQL_PASSWORD", "database_password")
}

type MysqlConfiguration struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

func (c MysqlConfiguration) Str() string {
	//"user:password@tcp(127.0.0.1:3306)/hello",
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		c.Username, c.Password, c.Host, c.Port, c.Database,
	)
}

func MakeMysqlConfig() MysqlConfiguration {
	return MysqlConfiguration{
		Host:     viper.GetString("MYSQL_HOST"),
		Port:     viper.GetString("MYSQL_PORT"),
		Database: viper.GetString("MYSQL_DATABASE"),
		Username: viper.GetString("MYSQL_USERNAME"),
		Password: viper.GetString("MYSQL_PASSWORD"),
	}
}
