package db

import "github.com/spf13/viper"

func init() {
	viper.SetDefault("DB_CONNECTION", "mysql")
}

type Config struct {
	Default     string
	Connections ConnectionsDbConfig
}

type ConnectionsDbConfig struct {
	Mysql   MysqlConfig
	Mariadb MariadbConfig
	Redis   RedisConfig
	Valkey  ValkeyConfig
}

func NewConfig() Config {
	return Config{
		Default: viper.GetString("DB_CONNECTION"),
		Connections: ConnectionsDbConfig{
			Mysql:   NewMysqlConfig(),
			Mariadb: NewMariadbConfig(),
			Redis:   NewRedisConfig(),
			Valkey:  NewValkeyConfig(),
		},
	}
}
