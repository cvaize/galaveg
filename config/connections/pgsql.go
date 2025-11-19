package connections

import (
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("PGSQL_HOST", "localhost")
	viper.SetDefault("PGSQL_PORT", 5432)
	viper.SetDefault("PGSQL_DATABASE", "database")
	viper.SetDefault("PGSQL_USERNAME", "database_user")
	viper.SetDefault("PGSQL_PASSWORD", "database_password")
	viper.SetDefault("PGSQL_PREFIX", "")
}

type PgsqlConfig struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
	Prefix   string
}

func (c *PgsqlConfig) DSN() string {
	return fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%d",
		c.Username, c.Password, c.Database, c.Host, c.Port,
	)
}

func NewPgsqlConfig() *PgsqlConfig {
	return &PgsqlConfig{
		Host:     viper.GetString("PGSQL_HOST"),
		Port:     viper.GetInt("PGSQL_PORT"),
		Database: viper.GetString("PGSQL_DATABASE"),
		Username: viper.GetString("PGSQL_USERNAME"),
		Password: viper.GetString("PGSQL_PASSWORD"),
		Prefix:   viper.GetString("PGSQL_PREFIX"),
	}
}
