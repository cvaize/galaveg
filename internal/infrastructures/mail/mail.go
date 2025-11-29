package mail

import (
	"galaveg/config"
	"galaveg/pkg/logger"
	"github.com/wneessen/go-mail"
	"strings"
)

type Mail = *mail.Client

// https://github.com/wneessen/go-mail/wiki/Getting-started#installation

func New(cfg *config.Config) (Mail, error) {
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

func Close(mail Mail) error {
	err := mail.Close()
	if err != nil {
		logger.Errorf("Mail Close() error: %s", err)
	}

	return err
}
