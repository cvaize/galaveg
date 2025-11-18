package config

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

// Add in the future transports: "sendmail", "mailgun", "ses", "ses-v2", "postmark", "resend", "log", "array", "failover", "roundrobin"

type SmtpMailerMailConfig struct {
	Transport  string
	Host       string
	Port       string
	Encryption string
	Username   string
	Password   string
}

type MailersMailConfig struct {
	Smtp SmtpMailerMailConfig
}

type FromMailConfig struct {
	Address string
	Name    string
}

type MailConfig struct {
	Default string
	Mailers MailersMailConfig
	From    FromMailConfig
}

func NewMailConfig() MailConfig {
	return MailConfig{
		Default: viper.GetString("MAIL_DEFAULT_MAILER"),
		Mailers: MailersMailConfig{
			Smtp: SmtpMailerMailConfig{
				Transport:  "smtp",
				Host:       viper.GetString("MAIL_HOST"),
				Port:       viper.GetString("MAIL_PORT"),
				Encryption: viper.GetString("MAIL_ENCRYPTION"),
				Username:   viper.GetString("MAIL_USERNAME"),
				Password:   viper.GetString("MAIL_PASSWORD"),
			},
		},
		From: FromMailConfig{
			Address: viper.GetString("MAIL_FROM_NAME"),
			Name:    viper.GetString("MAIL_FROM_ADDRESS"),
		},
	}
}
