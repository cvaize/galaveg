package config

import (
	"galaveg/config/app"
	"galaveg/config/connections"
	"galaveg/config/db"
	"galaveg/config/http"
	"galaveg/config/mail"
	"galaveg/config/session"
	"galaveg/utils/path"
	"github.com/spf13/viper"
	"path/filepath"
)

type Config struct {
	App         *app.Config
	Db          *db.Config
	Connections *connections.Config
	Http        *http.Config
	Mail        *mail.Config
	Session     *session.Config
}

func NewConfig(envFilename string) (*Config, error) {
	envPath := filepath.Join(path.FindModuleRoot(path.Cwd()), envFilename)
	viper.SetConfigFile(envPath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Config{
		App:         app.MustConfig(),
		Db:          db.MustConfig(),
		Connections: connections.MustConfig(),
		Http:        http.MustConfig(),
		Mail:        mail.MustConfig(),
		Session:     session.MustConfig(),
	}, nil
}

func MustConfig(envFilename string) *Config {
	c, e := NewConfig(envFilename)
	if e != nil {
		panic(e)
	}
	return c
}

func MustDefaultConfig() *Config {
	return MustConfig(".env")
}

func (c *Config) GetFolder(path string) string {
	return filepath.Join(c.App.Folder, path)
}
