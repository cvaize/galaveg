package mail

import (
	"galaveg/internal/modules/errors"
	"galaveg/internal/modules/mail/dto"
	libmail "github.com/wneessen/go-mail"
)

type Service struct {
	mail *libmail.Client
}

func NewService(mail *libmail.Client) (*Service, error) {
	return &Service{mail}, nil
}

// Send Synchronous sending of E-mail messages
func (s *Service) Send(message *dto.EmailMessage) *errors.Error {
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
		return errors.E500(e, "mail.Service.Send.DialAndSend", "")
	}
	return nil
}
