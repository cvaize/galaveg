package config

import (
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("REDIS_HOST", "localhost:6379")
	viper.SetDefault("REDIS_USERNAME", "")
	viper.SetDefault("REDIS_PASSWORD", "")
}

type RedisDbConfig struct {
	Host     string
	Username string
	Password string
}

func NewRedisDbConfig() RedisDbConfig {
	return RedisDbConfig{
		Host:     viper.GetString("REDIS_HOST"),
		Username: viper.GetString("REDIS_USERNAME"),
		Password: viper.GetString("REDIS_PASSWORD"),
	}
}
