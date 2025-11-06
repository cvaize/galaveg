package config

import (
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
	"time"
)

type Config struct {
	App AppConfig
	Db  DbConfig
}

func New(envPath string) Config {
	fmt.Println(envPath)
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

	return Config{
		App: NewAppConfig(),
		Db:  NewDbConfig(),
	}
}

func (c Config) GetFolder(path string) string {
	return filepath.Join(c.App.Folder, path)
}
