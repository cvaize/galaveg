package tests

import (
	"galaveg/config"
	"galaveg/utils"
)

func GetConfig() *config.Config {
	return utils.Must(config.New("tests.env"))
}
