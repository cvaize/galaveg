package main

import (
	"galaveg/cmd"
	"galaveg/config"
	"galaveg/utils/logger"
)

func main() {
	logger.SetLogLevel(config.Config.App.LogLevel)
	cmd.Execute()
}
