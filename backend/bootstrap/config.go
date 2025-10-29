package bootstrap

import (
	"galaveg/config"
	"galaveg/utils/logger"
	"github.com/spf13/viper"
)

func Config() {
	viper.SetDefault("APP_LOG_LEVEL", "info")

	if err := config.SetupConfig(); err != nil {
		logger.Fatalf("config SetupConfig() error: %s", err)
		panic(err)
	}

	logger.SetLogLevel(viper.GetString("APP_LOG_LEVEL"))
}
