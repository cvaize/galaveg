package auth

import (
	"galaveg/app/view/components/btn"
	"galaveg/app/view/components/field"
	"galaveg/app/view/layouts/auth/form"
	"galaveg/bootstrap/providers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type RegisterViewData struct {
	EmailValue            string
	EmailErrors           []string
	PasswordValue         string
	PasswordErrors        []string
	ConfirmPasswordValue  string
	ConfirmPasswordErrors []string
	Errors                []string
}

func NewRegister(c *gin.Context, ctx *providers.Context, s sessions.Session, data *RegisterViewData) (*View, error) {
	locale := ctx.S.LS.GetLocale(ctx.S.AS.Locale(c, nil))
	locales := ctx.S.LS.GetLocales()

	return &View{
		Lang:     locale.Code,
		DarkMode: ctx.S.AS.DarkMode(c),
		Title:    ctx.S.TS.T(locale.Code, "page.register.title"),
		Heading:  ctx.S.TS.T(locale.Code, "page.register.header"),
		Alerts:   ctx.S.AlS.Flashes(s),
		Locale:   locale,
		Locales:  locales,
		Form: form.View{
			Action: "/register",
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
					Label:      ctx.S.TS.T(locale.Code, "page.register.fields.email"),
					Type:       "email",
					Name:       "email",
					Value:      data.EmailValue,
					Errors:     data.EmailErrors,
					FieldClass: "admin-login__field",
					InputClass: "admin-login__field__input",
				},
				{
					Label:      ctx.S.TS.T(locale.Code, "page.register.fields.password"),
					Type:       "password",
					Name:       "password",
					Value:      data.PasswordValue,
					Errors:     data.PasswordErrors,
					FieldClass: "admin-login__field",
					InputClass: "admin-login__field__input",
				},
				{
					Label:      ctx.S.TS.T(locale.Code, "page.register.fields.confirm_password"),
					Type:       "password",
					Name:       "confirm_password",
					Value:      data.ConfirmPasswordValue,
					Errors:     data.ConfirmPasswordErrors,
					FieldClass: "admin-login__field",
					InputClass: "admin-login__field__input",
				},
			},
			Submit: &btn.View{
				Text: ctx.S.TS.T(locale.Code, "page.register.submit"),
			},
			ResetPassword: &btn.View{
				Href: "/reset-password",
				Text: ctx.S.TS.T(locale.Code, "page.register.reset_password"),
			},
			Login: &btn.View{
				Href: "/login",
				Text: ctx.S.TS.T(locale.Code, "page.register.login"),
			},
			Errors: data.Errors,
		},
		Back: &btn.View{Text: ctx.S.TS.T(locale.Code, "page.register.back"), Href: "/login"},
	}, nil
}
