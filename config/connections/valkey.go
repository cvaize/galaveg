package connections

import (
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("VALKEY_URL", "")
	viper.SetDefault("VALKEY_HOST", "localhost")
	viper.SetDefault("VALKEY_PORT", 6379)
	viper.SetDefault("VALKEY_USERNAME", "")
	viper.SetDefault("VALKEY_PASSWORD", "")
	viper.SetDefault("VALKEY_DB", 0)
	viper.SetDefault("VALKEY_MAX_RETRIES", 3)
}

type ValkeyConfig struct {
	Url        string
	Host       string
	Port       int
	Username   string
	Password   string
	Database   int
	MaxRetries int
}

func NewValkeyConfig() *ValkeyConfig {
	return &ValkeyConfig{
		Url:        viper.GetString("VALKEY_URL"),
		Host:       viper.GetString("VALKEY_HOST"),
		Port:       viper.GetInt("VALKEY_PORT"),
		Username:   viper.GetString("VALKEY_USERNAME"),
		Password:   viper.GetString("VALKEY_PASSWORD"),
		Database:   viper.GetInt("VALKEY_DB"),
		MaxRetries: viper.GetInt("VALKEY_MAX_RETRIES"),
	}
}
