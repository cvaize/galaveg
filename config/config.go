package config

import (
	"galaveg/config/app"
	"galaveg/config/auth"
	"galaveg/config/db"
	"galaveg/config/mail"
	"github.com/spf13/viper"
	"path/filepath"
	"time"
)

type Config struct {
	App  app.Config
	Db   db.Config
	Mail mail.Config
	Auth auth.Config
}

func New(envPath string) *Config {
	viper.SetConfigFile(envPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	//set timezone
	loc, err := time.LoadLocation(viper.GetString("APP_TIMEZONE"))
	if err != nil {
		panic(err)
	}
	time.Local = loc

	return &Config{
		App:  app.NewConfig(),
		Db:   db.NewConfig(),
		Mail: mail.NewConfig(),
		Auth: auth.NewConfig(),
	}
}

func (c *Config) GetFolder(path string) string {
	return filepath.Join(c.App.Folder, path)
}
