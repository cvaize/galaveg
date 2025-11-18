package db

import (
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("REDIS_URL", "")
	viper.SetDefault("REDIS_HOST", "localhost")
	viper.SetDefault("REDIS_PORT", 6379)
	viper.SetDefault("REDIS_USERNAME", "")
	viper.SetDefault("REDIS_PASSWORD", "")
	viper.SetDefault("REDIS_DB", 0)
	viper.SetDefault("REDIS_MAX_RETRIES", 3)
}

type RedisConfig struct {
	Url        string
	Host       string
	Port       int
	Username   string
	Password   string
	Database   int
	MaxRetries int
}

func NewRedisConfig() RedisConfig {
	return RedisConfig{
		Url:        viper.GetString("REDIS_URL"),
		Host:       viper.GetString("REDIS_HOST"),
		Port:       viper.GetInt("REDIS_PORT"),
		Username:   viper.GetString("REDIS_USERNAME"),
		Password:   viper.GetString("REDIS_PASSWORD"),
		Database:   viper.GetInt("REDIS_DB"),
		MaxRetries: viper.GetInt("REDIS_MAX_RETRIES"),
	}
}
