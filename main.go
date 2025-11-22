package main

import (
	"galaveg/cmd"
	"galaveg/utils/logger"
)

func main() {
	logger.SetLogLevel("info")
	cmd.Execute()
}
