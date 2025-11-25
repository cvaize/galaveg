package services

import (
	"galaveg/app/dto"
	"galaveg/config"
	"github.com/wneessen/go-mail"
)

type MailService struct {
	cfg    *config.Config
	client *mail.Client
}

func NewMailService(c *config.Config, client *mail.Client) (*MailService, error) {
	return &MailService{c, client}, nil
}

func MustMailService(c *config.Config, client *mail.Client) *MailService {
	s, e := NewMailService(c, client)
	if e != nil {
		panic(e)
	}
	return s
}

func (s *MailService) SendEmail(message *dto.EmailMessage) {

}
