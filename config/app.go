package config

import (
	"galaveg/utils/path"
	"github.com/samber/lo"
	"github.com/spf13/viper"
	"strings"
)

func init() {
	viper.SetDefault("APP_KEY", "")
	viper.SetDefault("APP_DEBUG", false)
	viper.SetDefault("APP_HOST", "0.0.0.0")
	viper.SetDefault("APP_URL", "http://localhost/")
	viper.SetDefault("APP_PORT", uint(8080))
	viper.SetDefault("APP_TIMEZONE", "UTC")
	viper.SetDefault("APP_LOG_LEVEL", "info")
	viper.SetDefault("APP_ALLOWED_HOSTS", "")
	viper.SetDefault("APP_FOLDER", path.FindModuleRoot(path.Cwd()))
	viper.SetDefault("APP_LOCALE", "en")
	viper.SetDefault("APP_LOCALE_COOKIE_KEY", "locale")
	viper.SetDefault("APP_DARK_MODE_COOKIE_KEY", "dark_mode")
}

type AppConfig struct {
	Key               string
	Debug             bool
	Host              string
	Url               string
	Port              uint
	Timezone          string
	LogLevel          string
	AllowedHosts      []string
	Folder            string
	Locale            string
	LocaleCookieKey   string
	DarkModeCookieKey string
}

func NewAppConfig() AppConfig {
	return AppConfig{
		Key:      viper.GetString("APP_KEY"),
		Debug:    viper.GetBool("APP_DEBUG"),
		Host:     viper.GetString("APP_HOST"),
		Url:      viper.GetString("APP_URL"),
		Port:     viper.GetUint("APP_PORT"),
		Timezone: viper.GetString("APP_TIMEZONE"),
		LogLevel: viper.GetString("APP_LOG_LEVEL"),
		AllowedHosts: lo.Filter(lo.Map(strings.Split(viper.GetString("APP_ALLOWED_HOSTS"), ","), func(s string, _ int) string {
			return strings.TrimSpace(s)
		}), func(s string, _ int) bool {
			return s != ""
		}),
		Folder:            viper.GetString("APP_FOLDER"),
		Locale:            viper.GetString("APP_LOCALE"),
		LocaleCookieKey:   viper.GetString("APP_LOCALE_COOKIE_KEY"),
		DarkModeCookieKey: viper.GetString("APP_DARK_MODE_COOKIE_KEY"),
	}
}
