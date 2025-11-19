package tests

import (
	"galaveg/config"
)

func GetConfig() *config.Config {
	return config.MustConfig("tests.env")
}
