package auth

import (
	"galaveg/app/view/components/btn"
	"galaveg/app/view/components/field"
	"galaveg/app/view/layouts/auth/form"
	"galaveg/bootstrap/providers"
	"github.com/gin-gonic/gin"
)

type LoginViewData struct {
	EmailValue     string
	EmailErrors    []string
	PasswordValue  string
	PasswordErrors []string
	Errors         []string
}

func NewLogin(c *gin.Context, ctx *providers.Context, data *LoginViewData) (*View, error) {
	locale := ctx.LS.GetLocale(ctx.AS.Locale(c, nil))
	locales := ctx.LS.GetLocales()

	return &View{
		Lang:     locale.Code,
		DarkMode: ctx.AS.DarkMode(c),
		Title:    ctx.TS.T(locale.Code, "page.login.title"),
		Heading:  ctx.TS.T(locale.Code, "page.login.header"),
		Alerts:   ctx.AS.Alerts(c),
		Locale:   locale,
		Locales:  locales,
		Form: form.View{
			Action: "/login",
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
					Label:      ctx.TS.T(locale.Code, "page.login.fields.email"),
					Type:       "email",
					Name:       "email",
					Value:      data.EmailValue,
					Errors:     data.EmailErrors,
					FieldClass: "admin-login__field",
					InputClass: "admin-login__field__input",
				},
				{
					Label:      ctx.TS.T(locale.Code, "page.login.fields.password"),
					Type:       "password",
					Name:       "password",
					Value:      data.PasswordValue,
					Errors:     data.PasswordErrors,
					FieldClass: "admin-login__field",
					InputClass: "admin-login__field__input",
				},
			},
			Submit: &btn.View{
				Text: ctx.TS.T(locale.Code, "page.login.submit"),
			},
			ResetPassword: &btn.View{
				Href: "/reset-password",
				Text: ctx.TS.T(locale.Code, "page.login.reset_password"),
			},
			Register: &btn.View{
				Href: "/register",
				Text: ctx.TS.T(locale.Code, "page.login.register"),
			},
			Errors: data.Errors,
		},
	}, nil
}
