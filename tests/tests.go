package tests

import (
	"galaveg/internal/config"
	"galaveg/pkg/utils"
)

func GetConfig() *config.Config {
	return utils.Must(config.New(".env.tests"))
}
