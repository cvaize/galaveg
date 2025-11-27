package providers

import (
	"galaveg/config"
	"github.com/wneessen/go-mail"
	"strings"
)

// https://github.com/wneessen/go-mail/wiki/Getting-started#installation

func NewMail(cfg *config.Config) (*mail.Client, error) {
	tls := mail.WithTLSPolicy(mail.NoTLS)
	if strings.ToLower(cfg.Mail.Encryption) == "tls" {
		tls = mail.WithTLSPolicy(mail.TLSMandatory)
	}

	client, err := mail.NewClient(cfg.Mail.Host, mail.WithPort(cfg.Mail.Port), mail.WithSMTPAuth(mail.SMTPAuthPlainNoEnc),
		mail.WithUsername(cfg.Mail.Username), mail.WithPassword(cfg.Mail.Password), tls, mail.WithLogAuthData())
	if err != nil {
		return client, err
	}
	return client, nil
}
