package services

import (
	"galaveg/app/dto"
	"galaveg/app/dto/mailables"
	libmail "github.com/wneessen/go-mail"
)

type MailService struct {
	mail *libmail.Client
	es   *ErrorService
}

func NewMailService() (*MailService, error) {
	return &MailService{}, nil
}

// Send Synchronous sending of E-mail messages
func (s *MailService) Send(message *mailables.EmailMessage) *dto.Error {
	m := libmail.NewMsg()
	m.FromMailAddress(message.Envelope.From)
	m.ToMailAddress(message.Envelope.To...)
	m.Subject(message.Envelope.Subject)
	// TODO: Explore and possibly implement the sending of HTML and TEXT and Markdown together, rather than just HTML or TEXT alone

	if message.Content.Html != "" {
		m.SetBodyString(libmail.TypeTextHTML, message.Content.Html)
	} else if message.Content.Text != "" {
		m.SetBodyString(libmail.TypeTextPlain, message.Content.Text)
	}

	if e := s.mail.DialAndSend(m); e != nil {
		return s.es.E500(e, "MailService.Send.DialAndSend", "")
	}
	return nil
}
