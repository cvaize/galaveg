package db

import (
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("MYSQL_DRIVER", "mysql")
	viper.SetDefault("MYSQL_HOST", "localhost")
	viper.SetDefault("MYSQL_PORT", 3306)
	viper.SetDefault("MYSQL_DATABASE", "database")
	viper.SetDefault("MYSQL_USERNAME", "database_user")
	viper.SetDefault("MYSQL_PASSWORD", "database_password")
	viper.SetDefault("MYSQL_PREFIX", "")
}

type MysqlConfig struct {
	Driver   string
	Host     string
	Port     int
	Database string
	Username string
	Password string
	Prefix   string
}

func (c *MysqlConfig) DSN() string {
	// TODO: Подумать над параметрами, на вроде multiStatements
	//"user:password@tcp(localhost:3306)/database?parseTime=true"
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		c.Username, c.Password, c.Host, c.Port, c.Database,
	)
}

func NewMysqlConfig() MysqlConfig {
	return MysqlConfig{
		Driver:   viper.GetString("MYSQL_DRIVER"),
		Host:     viper.GetString("MYSQL_HOST"),
		Port:     viper.GetInt("MYSQL_PORT"),
		Database: viper.GetString("MYSQL_DATABASE"),
		Username: viper.GetString("MYSQL_USERNAME"),
		Password: viper.GetString("MYSQL_PASSWORD"),
		Prefix:   viper.GetString("MYSQL_PREFIX"),
	}
}
