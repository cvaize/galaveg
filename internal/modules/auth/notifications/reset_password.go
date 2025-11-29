package notifications

import (
	"galaveg/internal/modules/errors"
	"galaveg/internal/modules/mail/dto"
	"net/mail"
)

type ResetPassword struct {
	Locale string
	Email  string
}

func NewResetPassword(locale string, email string) *ResetPassword {
	return &ResetPassword{locale, email}
}

func (s *ResetPassword) IsQueue() bool {
	return false
}

func (s *ResetPassword) Via() []string {
	return []string{"mail"}
}

func (s *ResetPassword) BuildEmailMessage() (*dto.EmailMessage, *errors.Error) {
	//resetPasswordLink, e1 := ctx.S.AuthS.CreateResetPasswordLink(s.Email)
	//if e1 != nil {
	//	return nil, ctx.S.ES.E500(e1, "ResetPasswordNotify.BuildEmailMessage.CreateResetPasswordLink", "")
	//}
	//
	//data, e2 := view.New(ctx, s.Locale, resetPasswordLink)
	//if e2 != nil {
	//	return nil, ctx.S.ES.E500(e2, "ResetPasswordNotify.BuildEmailMessage.NewData", "")
	//}
	//
	//html, e3 := ctx.S.TplS.Html(view.TEMPLATE, data)
	//if e3 != nil {
	//	return nil, ctx.S.ES.E500(e3, "ResetPasswordNotify.BuildEmailMessage.Html", "")
	//}
	//
	//return &EmailMessage{
	//	Envelope: &Envelope{
	//		Subject: ctx.S.TS.T(s.Locale, "mail.reset_password.subject"),
	//		From: &mail.Address{
	//			Address: ctx.Cfg.Mail.FromAddress,
	//			Name:    ctx.Cfg.Mail.FromName,
	//		},
	//		To: []*mail.Address{{Address: s.Email}},
	//	},
	//	Content: &Content{Html: html},
	//}, nil

	return &dto.EmailMessage{
		Envelope: &dto.Envelope{
			Subject: "test",
			From: &mail.Address{
				Address: "test@test.ru",
				Name:    "test",
			},
			To: []*mail.Address{{Address: s.Email}},
		},
		Content: &dto.Content{Html: "test"},
	}, nil
}
