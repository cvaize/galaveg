package auth

import (
	"galaveg/app/view/components/btn"
	"galaveg/app/view/components/field"
	"galaveg/app/view/layouts/auth/form"
	"galaveg/bootstrap/providers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type ResetPasswordViewData struct {
	EmailValue  string
	EmailErrors []string
	Errors      []string
}

func NewResetPassword(c *gin.Context, ctx *providers.Context, s sessions.Session, data *ResetPasswordViewData) (*View, error) {
	locale := ctx.S.LS.GetLocale(ctx.S.AS.Locale(c, nil))
	locales := ctx.S.LS.GetLocales()

	return &View{
		Lang:     locale.Code,
		DarkMode: ctx.S.AS.DarkMode(c),
		Title:    ctx.S.TS.T(locale.Code, "page.reset_password.title"),
		Heading:  ctx.S.TS.T(locale.Code, "page.reset_password.header"),
		Alerts:   ctx.S.AlS.Flashes(s),
		Locale:   locale,
		Locales:  locales,
		Form: form.View{
			Action: "/reset-password",
			Method: "post",
			Fields: []field.View{
				{
					Type:       "hidden",
					Name:       "csrf",
					Value:      "",
					FieldClass: "admin-login__field",
					InputClass: "admin-login__field__input",
				},
				{
					Label:      ctx.S.TS.T(locale.Code, "page.reset_password.fields.email"),
					Type:       "email",
					Name:       "email",
					Value:      data.EmailValue,
					Errors:     data.EmailErrors,
					FieldClass: "admin-login__field",
					InputClass: "admin-login__field__input",
				},
			},
			Submit: &btn.View{
				Text: ctx.S.TS.T(locale.Code, "page.reset_password.submit"),
			},
			Errors: data.Errors,
			Text:   ctx.S.TS.T(locale.Code, "page.reset_password.text"),
		},
		Back: &btn.View{Text: ctx.S.TS.T(locale.Code, "page.reset_password.back"), Href: "/login"},
	}, nil
}
