package reset_password

import (
	"galaveg/internal/modules/errors"
	"galaveg/internal/modules/mail/dto"
	"galaveg/internal/modules/notifications"
	view "galaveg/internal/modules/view/layouts/email/reset_password"
	"net/mail"
)

type Notification struct {
	locale string
	email  string
	link   string
}

func NewNotification(locale string, email, link string) *Notification {
	return &Notification{locale, email, link}
}

func (s *Notification) IsQueue(ctx *notifications.BuildNotificationContext) bool {
	return false
}

func (s *Notification) Via(ctx *notifications.BuildNotificationContext) []string {
	return []string{"mail"}
}

func (s *Notification) BuildEmailMessage(ctx *notifications.BuildNotificationContext) (*dto.EmailMessage, *errors.Error) {
	data, e2 := view.New(ctx.AppService, ctx.TranslatorService, s.locale, s.link)
	if e2 != nil {
		return nil, errors.E500(e2, "auth.notifications.Notification.BuildEmailMessage.NewData", "")
	}

	html, e3 := ctx.TemplateService.Html(view.TEMPLATE, data)
	if e3 != nil {
		return nil, errors.E500(e3, "auth.notifications.Notification.BuildEmailMessage.Html", "")
	}

	return &dto.EmailMessage{
		Envelope: &dto.Envelope{
			Subject: ctx.TranslatorService.T(s.locale, "mail.reset_password.subject"),
			From: &mail.Address{
				Address: ctx.Cfg.Mail.FromAddress,
				Name:    ctx.Cfg.Mail.FromName,
			},
			To: []*mail.Address{{Address: s.email}},
		},
		Content: &dto.Content{Html: html},
	}, nil
}
