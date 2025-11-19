package http

import (
	"github.com/samber/lo"
	"github.com/spf13/viper"
	"strings"
)

func init() {
	viper.SetDefault("HTTP_HOST", "localhost")
	viper.SetDefault("HTTP_PORT", 80)
	// HTTP_SCHEMA = https, http
	viper.SetDefault("HTTP_SCHEMA", "http")
	// HTTP_ALLOWED_HOSTS = localhost,0.0.0.0,example.com
	viper.SetDefault("HTTP_ALLOWED_HOSTS", "")
}

type Config struct {
	Host         string
	Port         int
	Schema       string
	AllowedHosts []string
}

func NewConfig() (*Config, error) {
	return &Config{
		Host:   viper.GetString("HTTP_HOST"),
		Port:   viper.GetInt("HTTP_PORT"),
		Schema: viper.GetString("HTTP_SCHEMA"),
		AllowedHosts: lo.Filter(lo.Map(strings.Split(viper.GetString("HTTP_ALLOWED_HOSTS"), ","), func(s string, _ int) string {
			return strings.TrimSpace(s)
		}), func(s string, _ int) bool {
			return s != ""
		}),
	}, nil
}

func MustConfig() *Config {
	c, e := NewConfig()
	if e != nil {
		panic(e)
	}
	return c
}
