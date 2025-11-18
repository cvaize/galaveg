package config

import (
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("AUTH_COOKIE_KEY", "session")
}

type AuthConfig struct {
	SessionStoreUserKey string
	CookieKey           string
}

func NewAuthConfig() AuthConfig {
	return AuthConfig{
		SessionStoreUserKey: "user_id",
		CookieKey:           viper.GetString("AUTH_COOKIE_KEY"),
	}
}
