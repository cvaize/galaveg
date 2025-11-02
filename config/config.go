package config

import (
	"github.com/spf13/viper"
	"time"
)

type Configuration struct {
	App AppConfiguration
	Db  DbConfiguration
}

// Config - Глобальный конфиг. Заполняется при запуске.
// Изменять его во время исполнения запрещено, иначе будет гонка данных.
var Config Configuration

func init() {
	Config = Make()
}

func Make() Configuration {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	//set timezone
	loc, err := time.LoadLocation(viper.GetString("APP_TIMEZONE"))
	if err != nil {
		panic(err)
	}
	time.Local = loc

	return Configuration{
		App: MakeAppConfig(),
		Db:  MakeDbConfig(),
	}
}
