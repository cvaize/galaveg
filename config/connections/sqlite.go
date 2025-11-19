package connections

import (
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("SQLITE_DSN", "./storage/test.db")
}

type SqliteConfig struct {
	Dsn string
}

func NewSqliteConfig() *SqliteConfig {
	return &SqliteConfig{
		Dsn: viper.GetString("SQLITE_DSN"),
	}
}
