package auth

import (
	"galaveg/app/view/components/btn"
	"galaveg/app/view/components/field"
	"galaveg/app/view/layouts/auth/form"
	"galaveg/bootstrap/providers"
	"github.com/gin-gonic/gin"
)

type ResetPasswordViewData struct {
	EmailValue  string
	EmailErrors []string
	Errors      []string
}

func NewResetPassword(c *gin.Context, ctx *providers.Context, data *ResetPasswordViewData) (*View, error) {
	locale := ctx.LS.GetLocale(ctx.AS.Locale(c, nil))
	locales := ctx.LS.GetLocales()

	return &View{
		Lang:     locale.Code,
		DarkMode: ctx.AS.DarkMode(c),
		Title:    ctx.TS.T(locale.Code, "page.reset_password.title"),
		Heading:  ctx.TS.T(locale.Code, "page.reset_password.header"),
		Alerts:   ctx.AS.Alerts(c),
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
					Label:      ctx.TS.T(locale.Code, "page.reset_password.fields.email"),
					Type:       "email",
					Name:       "email",
					Value:      data.EmailValue,
					Errors:     data.EmailErrors,
					FieldClass: "admin-login__field",
					InputClass: "admin-login__field__input",
				},
			},
			Submit: &btn.View{
				Text: ctx.TS.T(locale.Code, "page.reset_password.submit"),
			},
			Errors: data.Errors,
			Text:   ctx.TS.T(locale.Code, "page.reset_password.text"),
		},
		Back: &btn.View{Text: ctx.TS.T(locale.Code, "page.reset_password.back"), Href: "/login"},
	}, nil
}
