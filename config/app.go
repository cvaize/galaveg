package config

import (
	"github.com/samber/lo"
	"github.com/spf13/viper"
	"strings"
)

func init() {
	viper.SetDefault("APP_DEBUG", false)
	viper.SetDefault("APP_HOST", "0.0.0.0")
	viper.SetDefault("APP_PORT", uint(8080))
	viper.SetDefault("APP_TIMEZONE", "UTC")
	viper.SetDefault("APP_LOG_LEVEL", "info")
	viper.SetDefault("APP_ALLOWED_HOSTS", "")
}

type AppConfiguration struct {
	Debug        bool
	Host         string
	Port         uint
	Timezone     string
	LogLevel     string
	AllowedHosts []string
}

func MakeAppConfig() AppConfiguration {
	return AppConfiguration{
		Debug:    viper.GetBool("APP_DEBUG"),
		Host:     viper.GetString("APP_HOST"),
		Port:     viper.GetUint("APP_PORT"),
		Timezone: viper.GetString("APP_TIMEZONE"),
		LogLevel: viper.GetString("APP_LOG_LEVEL"),
		AllowedHosts: lo.Filter(lo.Map(strings.Split(viper.GetString("APP_ALLOWED_HOSTS"), ","), func(s string, _ int) string {
			return strings.TrimSpace(s)
		}), func(s string, _ int) bool {
			return s != ""
		}),
	}
}
