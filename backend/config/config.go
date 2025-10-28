package config

import (
	"galaveg/utils/logger"
	"github.com/spf13/viper"
	"time"
)

//type Configuration struct {
//	Server   ServerConfiguration
//	Database DatabaseConfiguration
//}

// SetupConfig configuration
func SetupConfig() error {
	//var configuration *Configuration
	viper.SetDefault("APP_TIMEZONE", "UTC")

	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		logger.Errorf("Error to reading config file, %s", err)
		return err
	}

	//err := viper.Unmarshal(&configuration)
	//if err != nil {
	//	utils.Errorf("error to decode, %v", err)
	//	return err
	//}

	//set timezone
	loc, err := time.LoadLocation(viper.GetString("APP_TIMEZONE"))
	if err != nil {
		logger.Errorf("Error to time.LoadLocation: %s", err)
		return err
	}
	time.Local = loc

	return nil
}
