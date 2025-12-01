package auth

import (
	"galaveg/internal/modules/app"
	"galaveg/internal/modules/translator"
	"galaveg/internal/modules/view/components/btn"
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

func New(as app.Service, ts translator.Service, lang, resetPasswordLink string) (*View, error) {
	return &View{
		Lang:        lang,
		Title:       ts.T(lang, "mail.reset_password.title"),
		Description: ts.T(lang, "mail.reset_password.description"),
		SiteName:    ts.T(lang, "mail.reset_password.site_name"),
		SiteURL:     as.Url(),
		LogoSrc:     as.LogoSrc(),
		Header:      ts.T(lang, "mail.reset_password.header"),
		SiteDomain:  as.CloneUrl().Host,
		Button: &btn.View{
			Text: ts.T(lang, "mail.reset_password.button"),
			Href: resetPasswordLink,
		},
	}, nil
}
