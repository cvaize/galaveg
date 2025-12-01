package notifications

import (
	"galaveg/internal/modules/auth/actions"
	"galaveg/internal/modules/errors"
	"galaveg/internal/modules/mail/dto"
	"galaveg/internal/modules/notifications"
	view "galaveg/internal/modules/view/layouts/email/reset_password"
	"net/mail"
)

type ResetPassword struct {
	Locale string
	Email  string
}

func NewResetPassword(locale string, email string) *ResetPassword {
	return &ResetPassword{locale, email}
}

func (s *ResetPassword) IsQueue(ctx *notifications.BuildNotificationContext) bool {
	return false
}

func (s *ResetPassword) Via(ctx *notifications.BuildNotificationContext) []string {
	return []string{"mail"}
}

func (s *ResetPassword) BuildEmailMessage(ctx *notifications.BuildNotificationContext) (*dto.EmailMessage, *errors.Error) {
	resetPasswordLink, e1 := actions.CreateResetPasswordLink(ctx.AppService, s.Email)
	if e1 != nil {
		return nil, errors.E500(e1, "auth.notifications.ResetPassword.BuildEmailMessage.CreateResetPasswordLink", "")
	}

	data, e2 := view.New(ctx.AppService, ctx.TranslatorService, s.Locale, resetPasswordLink)
	if e2 != nil {
		return nil, errors.E500(e2, "auth.notifications.ResetPassword.BuildEmailMessage.NewData", "")
	}

	html, e3 := ctx.TemplateService.Html(view.TEMPLATE, data)
	if e3 != nil {
		return nil, errors.E500(e3, "auth.notifications.ResetPassword.BuildEmailMessage.Html", "")
	}

	return &dto.EmailMessage{
		Envelope: &dto.Envelope{
			Subject: ctx.TranslatorService.T(s.Locale, "mail.reset_password.subject"),
			From: &mail.Address{
				Address: ctx.Cfg.Mail.FromAddress,
				Name:    ctx.Cfg.Mail.FromName,
			},
			To: []*mail.Address{{Address: s.Email}},
		},
		Content: &dto.Content{Html: html},
	}, nil
}
