package mail

import (
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("MAIL_DEFAULT_MAILER", "smtp")
	viper.SetDefault("MAIL_HOST", "fake-smtp-server")
	viper.SetDefault("MAIL_PORT", "8025")
	viper.SetDefault("MAIL_ENCRYPTION", "")
	viper.SetDefault("MAIL_USERNAME", "test")
	viper.SetDefault("MAIL_PASSWORD", "test")
	viper.SetDefault("MAIL_FROM_NAME", "test")
	viper.SetDefault("MAIL_FROM_ADDRESS", "fake@fake.com")
}

type SmtpMailerConfig struct {
	Transport  string
	Host       string
	Port       string
	Encryption string
	Username   string
	Password   string
}

type MailersConfig struct {
	Smtp SmtpMailerConfig
}

type FromConfig struct {
	Address string
	Name    string
}

type Config struct {
	Default string
	Mailers MailersConfig
	From    FromConfig
}

func NewConfig() (*Config, error) {
	return &Config{
		Default: viper.GetString("MAIL_DEFAULT_MAILER"),
		Mailers: MailersConfig{
			Smtp: SmtpMailerConfig{
				Transport:  "smtp",
				Host:       viper.GetString("MAIL_HOST"),
				Port:       viper.GetString("MAIL_PORT"),
				Encryption: viper.GetString("MAIL_ENCRYPTION"),
				Username:   viper.GetString("MAIL_USERNAME"),
				Password:   viper.GetString("MAIL_PASSWORD"),
			},
		},
		From: FromConfig{
			Address: viper.GetString("MAIL_FROM_NAME"),
			Name:    viper.GetString("MAIL_FROM_ADDRESS"),
		},
	}, nil
}

func MustConfig() *Config {
	c, e := NewConfig()
	if e != nil {
		panic(e)
	}
	return c
}
