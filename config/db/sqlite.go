package db

import (
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("SQLITE_DRIVER", "sqlite3")
	viper.SetDefault("SQLITE_DSN", "./storage/test.db")
}

type SqliteConfig struct {
	Driver string
	Dsn    string
}

func NewSqliteConfig() SqliteConfig {
	return SqliteConfig{
		Driver: viper.GetString("SQLITE_DRIVER"),
		Dsn:    viper.GetString("SQLITE_DSN"),
	}
}
