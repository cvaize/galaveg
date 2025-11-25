package services

import (
	"galaveg/app/dto"
	"galaveg/app/dto/mailables"
	"galaveg/config"
	"galaveg/utils/logger"
	libmail "github.com/wneessen/go-mail"
	"net/mail"
)

type MailService struct {
	cfg    *config.Config
	client *libmail.Client
}

// TODO: Make messages sent by cron through the message buffer.

func NewMailService(c *config.Config, client *libmail.Client) (*MailService, error) {
	return &MailService{c, client}, nil
}

func MustMailService(c *config.Config, client *libmail.Client) *MailService {
	s, e := NewMailService(c, client)
	if e != nil {
		panic(e)
	}
	return s
}

func (s *MailService) NewSimpleEmailMessage(to, subject, html, txt string) *dto.EmailMessage {
	return &dto.EmailMessage{
		Envelope: &mailables.Envelope{
			Subject: subject,
			From: &mail.Address{
				Address: s.cfg.Mail.FromAddress,
				Name:    s.cfg.Mail.FromName,
			},
			To: []*mail.Address{{Address: to}},
		},
		Content: &mailables.Content{HtmlString: html, Text: txt},
	}
}

func (s *MailService) SendEmail(message *dto.EmailMessage) error {
	m := libmail.NewMsg()
	m.FromMailAddress(message.Envelope.From)
	m.ToMailAddress(message.Envelope.To...)
	m.Subject(message.Envelope.Subject)
	// TODO: Explore and possibly implement the sending of HTML and TEXT and Markdown together, rather than just HTML or TEXT alone
	// TODO: Add TEXT and HTML rendering. The SetBodyHTMLTemplate and SetBodyTextTemplate methods.

	if message.Content.HtmlString != "" {
		m.SetBodyString(libmail.TypeTextHTML, message.Content.HtmlString)
	} else if message.Content.Text != "" {
		m.SetBodyString(libmail.TypeTextPlain, message.Content.Text)
	}

	if err := s.client.DialAndSend(m); err != nil {
		logger.Errorf("(500) MailService.SendEmail.DialAndSend: %v", err)
		return err
	}
	return nil
}

func (s *MailService) SendSimpleEmailMessage(to, subject, html, txt string) error {
	message := s.NewSimpleEmailMessage(to, subject, html, txt)
	return s.SendEmail(message)
}
