package notifications

import (
	"galaveg/app/dto"
	"galaveg/app/dto/mailables"
	view "galaveg/app/view/layouts/email/reset_password"
	"galaveg/bootstrap/providers"
	"net/mail"
)

type ResetPassword struct {
	Locale string
	Email  string
}

func NewResetPassword(locale string, email string) *ResetPassword {
	return &ResetPassword{locale, email}
}

func (s *ResetPassword) IsQueue(ctx *providers.Context) bool {
	return false
}

func (s *ResetPassword) Via(ctx *providers.Context) []string {
	return []string{"mail"}
}

func (s *ResetPassword) BuildEmailMessage(ctx *providers.Context) (*mailables.EmailMessage, *dto.Error) {
	resetPasswordLink, e1 := ctx.S.AuthS.CreateResetPasswordLink(s.Email)
	if e1 != nil {
		return nil, ctx.S.ES.E500(e1, "ResetPasswordNotify.BuildEmailMessage.CreateResetPasswordLink", "")
	}

	data, e2 := view.New(ctx, s.Locale, resetPasswordLink)
	if e2 != nil {
		return nil, ctx.S.ES.E500(e2, "ResetPasswordNotify.BuildEmailMessage.NewData", "")
	}

	html, e3 := ctx.S.TplS.Html(view.TEMPLATE, data)
	if e3 != nil {
		return nil, ctx.S.ES.E500(e3, "ResetPasswordNotify.BuildEmailMessage.Html", "")
	}

	return &mailables.EmailMessage{
		Envelope: &mailables.Envelope{
			Subject: ctx.S.TS.T(s.Locale, "mail.reset_password.subject"),
			From: &mail.Address{
				Address: ctx.Cfg.Mail.FromAddress,
				Name:    ctx.Cfg.Mail.FromName,
			},
			To: []*mail.Address{{Address: s.Email}},
		},
		Content: &mailables.Content{Html: html},
	}, nil
}
