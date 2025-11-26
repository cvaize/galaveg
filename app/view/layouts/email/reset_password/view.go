package auth

import (
	"galaveg/app/view/components/btn"
	"galaveg/bootstrap/providers"
	"github.com/gin-gonic/gin"
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

func New(c *gin.Context, ctx *providers.Context, data *ViewData) (*View, error) {
	locale := ctx.LS.GetLocale(ctx.AS.Locale(c, nil))

	return &View{
		Lang:        locale.Code,
		Title:       ctx.TS.T(locale.Code, "mail.reset_password.title"),
		Description: ctx.TS.T(locale.Code, "mail.reset_password.description"),
		SiteName:    ctx.TS.T(locale.Code, "mail.reset_password.site_name"),
		SiteURL:     ctx.AS.Url(),
		LogoSrc:     ctx.AS.LogoSrc(),
		Header:      ctx.TS.T(locale.Code, "mail.reset_password.header"),
		SiteDomain:  ctx.AS.CloneUrl().Host,
		Button: &btn.View{
			Text: ctx.TS.T(locale.Code, "mail.reset_password.button"),
			Href: data.ResetPasswordLink,
		},
	}, nil
}
