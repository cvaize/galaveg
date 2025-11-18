package main

import (
	"galaveg/cmd"
	"galaveg/utils/logger"
	"github.com/spf13/viper"
)

func main() {
	viper.SetDefault("APP_LOG_LEVEL", "info")
	logger.SetLogLevel(viper.GetString("APP_LOG_LEVEL"))
	cmd.Execute()
}
