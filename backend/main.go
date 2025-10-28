package main

import (
	"galaveg/app/console/commands/migrate"
	"galaveg/bootstrap"
	"galaveg/config"
	"galaveg/utils/logger"
	"github.com/samber/lo"
	"github.com/spf13/viper"
	"os"
)

func main() {
	viper.SetDefault("APP_LOG_LEVEL", "info")

	if err := config.SetupConfig(); err != nil {
		logger.Fatalf("config SetupConfig() error: %s", err)
		panic(err)
	}

	logger.SetLogLevel(viper.GetString("APP_LOG_LEVEL"))

	switch lo.NthOr(os.Args, 1, "") {
	case "server":
		bootstrap.Server()
		break
	case "migrate":
		switch lo.NthOr(os.Args, 2, "") {
		case "up":
			migrate.Up()
			break
		case "down":
			migrate.Down()
			break
		}
		break
	}
}
