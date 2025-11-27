package auth

import (
	"galaveg/app/view/components/btn"
	"galaveg/bootstrap/providers"
)

const TEMPLATE = "layouts/email/reset_password"

type View struct {
	Lang        string
	Title       string
	Description string
	SiteURL     string
	LogoSrc     string
	SiteName    string
	Header      string
	SiteDomain  string
	Button      *btn.View
}

type ViewData struct {
	ResetPasswordLink string
}

func New(ctx *providers.Context, lang, resetPasswordLink string) (*View, error) {
	return &View{
		Lang:        lang,
		Title:       ctx.S.TS.T(lang, "mail.reset_password.title"),
		Description: ctx.S.TS.T(lang, "mail.reset_password.description"),
		SiteName:    ctx.S.TS.T(lang, "mail.reset_password.site_name"),
		SiteURL:     ctx.S.AS.Url(),
		LogoSrc:     ctx.S.AS.LogoSrc(),
		Header:      ctx.S.TS.T(lang, "mail.reset_password.header"),
		SiteDomain:  ctx.S.AS.CloneUrl().Host,
		Button: &btn.View{
			Text: ctx.S.TS.T(lang, "mail.reset_password.button"),
			Href: resetPasswordLink,
		},
	}, nil
}
