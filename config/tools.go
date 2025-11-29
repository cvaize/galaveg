package config

import (
	"fmt"
	"galaveg/pkg/logger"
	"galaveg/pkg/utils"
	"github.com/joho/godotenv"
	"github.com/samber/lo"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func New(envFilename string) (*Config, error) {
	v := viper.New()
	c := Config{}

	envPath := filepath.Join(utils.FindModuleRoot(utils.Cwd()), envFilename)
	if err := godotenv.Load(envPath); err != nil {
		if !os.IsNotExist(err) {
			log.Printf("Warning: .env file not loaded: %v", err)
		}
	}

	configPath := filepath.Join(utils.FindModuleRoot(utils.Cwd()), "config", "default.yaml")

	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.ReadInConfig(); err != nil {
		logger.Fatalf("Fatal error config file: %v", err)
		return nil, err
	}

	if err := v.Unmarshal(&c); err != nil {
		log.Fatalf("Unable to decode config into struct: %v", err)
	}

	beforeReturn(&c)

	if c.App.Debug {
		fmt.Printf("Config loaded: %+v\n", c)
	}

	return &c, nil
}

func Default() (*Config, error) {
	return New(".env")
}

func MustDefault() *Config {
	return utils.Must(Default())
}

func strDefault(v string, d string) string {
	if v == "" {
		return d
	}
	return v
}

func strSliceFilter(v []string) []string {
	return lo.Filter(lo.Map(v, func(s string, _ int) string {
		return strings.TrimSpace(s)
	}), func(s string, _ int) bool {
		return s != ""
	})
}

func (c *Config) GetFolder(path string) string {
	return filepath.Join(c.App.Folder, path)
}
