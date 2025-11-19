package connections

import (
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("MONGODB_URI", "")
	viper.SetDefault("MONGODB_HOST", "localhost")
	viper.SetDefault("MONGODB_PORT", 27017)
	viper.SetDefault("MONGODB_DATABASE", "database")
}

type MongodbConfig struct {
	Uri      string
	Host     string
	Port     int
	Database string
}

func NewMongodbConfig() *MongodbConfig {
	return &MongodbConfig{
		Uri:      viper.GetString("MONGODB_URI"),
		Host:     viper.GetString("MONGODB_HOST"),
		Port:     viper.GetInt("MONGODB_PORT"),
		Database: viper.GetString("MONGODB_DATABASE"),
	}
}
