package main

import (
	"galaveg/cmd"
	"galaveg/config"
	"galaveg/utils/logger"
)

func main() {
	// TODO: Поместить cobra-cli и migrate в проект
	logger.SetLogLevel(config.Config.App.LogLevel)
	cmd.Execute()
}
