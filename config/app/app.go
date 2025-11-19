package app

import (
	"galaveg/utils/path"
	"github.com/samber/lo"
	"github.com/spf13/viper"
	"strings"
	"time"
)

func init() {
	viper.SetDefault("APP_KEY", "")
	viper.SetDefault("APP_PREVIOUS_KEYS", "")
	viper.SetDefault("APP_URL", "http://localhost/")
	viper.SetDefault("APP_DEBUG", false)
	viper.SetDefault("APP_TIMEZONE", "UTC")
	// panic, fatal, error, warn, info, debug, trace
	viper.SetDefault("APP_LOG_LEVEL", "info")
	viper.SetDefault("APP_FOLDER", path.FindModuleRoot(path.Cwd()))
	viper.SetDefault("APP_LOCALE", "en")
	viper.SetDefault("APP_LOCALE_COOKIE_KEY", "locale")
	viper.SetDefault("APP_DARK_MODE_COOKIE_KEY", "dark_mode")
}

type Config struct {
	Key               string
	PreviousKeys      []string
	Url               string
	Debug             bool
	Timezone          string
	LogLevel          string
	Folder            string
	Locale            string
	LocaleCookieKey   string
	DarkModeCookieKey string
}

func NewConfig() (*Config, error) {
	//set timezone
	loc, err := time.LoadLocation(viper.GetString("APP_TIMEZONE"))
	if err != nil {
		return nil, err
	}
	time.Local = loc

	return &Config{
		Key: viper.GetString("APP_KEY"),
		PreviousKeys: lo.Filter(lo.Map(strings.Split(viper.GetString("APP_PREVIOUS_KEYS"), ","), func(s string, _ int) string {
			return strings.TrimSpace(s)
		}), func(s string, _ int) bool {
			return s != ""
		}),
		Url:               viper.GetString("APP_URL"),
		Debug:             viper.GetBool("APP_DEBUG"),
		Timezone:          viper.GetString("APP_TIMEZONE"),
		LogLevel:          viper.GetString("APP_LOG_LEVEL"),
		Folder:            viper.GetString("APP_FOLDER"),
		Locale:            viper.GetString("APP_LOCALE"),
		LocaleCookieKey:   viper.GetString("APP_LOCALE_COOKIE_KEY"),
		DarkModeCookieKey: viper.GetString("APP_DARK_MODE_COOKIE_KEY"),
	}, nil
}

func MustConfig() *Config {
	c, e := NewConfig()
	if e != nil {
		panic(e)
	}
	return c
}
