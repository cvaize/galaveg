package auth

import (
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("AUTH_COOKIE_KEY", "session")
}

type Config struct {
	SessionStoreUserKey string
	CookieKey           string
}

func NewConfig() Config {
	return Config{
		SessionStoreUserKey: "user_id",
		CookieKey:           viper.GetString("AUTH_COOKIE_KEY"),
	}
}
