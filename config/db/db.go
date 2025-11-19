package db

import "github.com/spf13/viper"

func init() {
	viper.SetDefault("DB_PREFIX", "")
}

type Config struct {
	Prefix string
}

func NewConfig() (*Config, error) {
	return &Config{
		Prefix: viper.GetString("DB_PREFIX"),
	}, nil
}

func MustConfig() *Config {
	c, e := NewConfig()
	if e != nil {
		panic(e)
	}
	return c
}
