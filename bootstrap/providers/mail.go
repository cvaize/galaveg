package providers

import (
	"galaveg/config"
	"github.com/wneessen/go-mail"
)

// https://github.com/wneessen/go-mail/wiki/Getting-started#installation

func NewMail(cfg *config.Config) (*mail.Client, error) {
	client, err := mail.NewClient(cfg.Mail.Host, mail.WithPort(cfg.Mail.Port), mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(cfg.Mail.Username), mail.WithPassword(cfg.Mail.Password))
	if err != nil {
		return client, err
	}
	return client, nil
}

func MustMail(cfg *config.Config) *mail.Client {
	client, e := NewMail(cfg)
	if e != nil {
		panic(e)
	}
	return client
}
