package auth

import (
	"galaveg/app/view/components/btn"
	"galaveg/app/view/components/field"
	"galaveg/app/view/layouts/auth/form"
	"galaveg/bootstrap/providers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginViewData struct {
	EmailValue     string
	EmailErrors    []string
	PasswordValue  string
	PasswordErrors []string
	Errors         []string
}

func NewLogin(c *gin.Context, ctx *providers.Context, s sessions.Session, data *LoginViewData) (*View, error) {
	locale := ctx.S.LS.GetLocale(ctx.S.AS.Locale(c, nil))
	locales := ctx.S.LS.GetLocales()

	return &View{
		Lang:     locale.Code,
		DarkMode: ctx.S.AS.DarkMode(c),
		Title:    ctx.S.TS.T(locale.Code, "page.login.title"),
		Heading:  ctx.S.TS.T(locale.Code, "page.login.header"),
		Alerts:   ctx.S.AlS.Flashes(s),
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
					Label:      ctx.S.TS.T(locale.Code, "page.login.fields.email"),
					Type:       "email",
					Name:       "email",
					Value:      data.EmailValue,
					Errors:     data.EmailErrors,
					FieldClass: "admin-login__field",
					InputClass: "admin-login__field__input",
				},
				{
					Label:      ctx.S.TS.T(locale.Code, "page.login.fields.password"),
					Type:       "password",
					Name:       "password",
					Value:      data.PasswordValue,
					Errors:     data.PasswordErrors,
					FieldClass: "admin-login__field",
					InputClass: "admin-login__field__input",
				},
			},
			Submit: &btn.View{
				Text: ctx.S.TS.T(locale.Code, "page.login.submit"),
			},
			ResetPassword: &btn.View{
				Href: "/reset-password",
				Text: ctx.S.TS.T(locale.Code, "page.login.reset_password"),
			},
			Register: &btn.View{
				Href: "/register",
				Text: ctx.S.TS.T(locale.Code, "page.login.register"),
			},
			Errors: data.Errors,
		},
	}, nil
}
