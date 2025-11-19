package session

import (
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("SESSION_STORE_CONNECTION", "redis")
	viper.SetDefault("SESSION_STORE_USER_KEY", "user_id")
	viper.SetDefault("SESSION_COOKIE", "session_id")
}

type Config struct {
	StoreConnection string
	StoreUserKey    string
	Cookie          string
}

func NewConfig() (*Config, error) {
	return &Config{
		StoreConnection: viper.GetString("SESSION_STORE_CONNECTION"),
		StoreUserKey:    viper.GetString("SESSION_STORE_USER_KEY"),
		Cookie:          viper.GetString("SESSION_COOKIE"),
	}, nil
}

func MustConfig() *Config {
	c, e := NewConfig()
	if e != nil {
		panic(e)
	}
	return c
}
