package tests

import (
	"galaveg/config"
	"github.com/spf13/viper"
	"path/filepath"
)

func GetEnv() string {
	return filepath.Join(viper.GetString("APP_FOLDER"), "tests.env")
}

func GetConfig() config.Config {
	return config.New(GetEnv())
}
