package main

import (
	"galaveg/bootstrap/singleton"
	"galaveg/cmd"
	"galaveg/utils/logger"
)

func main() {
	logger.SetLogLevel(singleton.C.App.LogLevel)
	cmd.Execute()
}
