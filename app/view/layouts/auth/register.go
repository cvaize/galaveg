package auth

import (
	"galaveg/app/view/components/btn"
	"galaveg/app/view/components/field"
	"galaveg/app/view/layouts/auth/form"
	"galaveg/bootstrap/providers"
	"github.com/gin-gonic/gin"
)

func NewRegister(c *gin.Context, ctx *providers.Context) (*View, error) {
	locale := ctx.LS.GetLocale(ctx.AS.Locale(c, nil))
	locales := ctx.LS.GetLocales()

	return &View{
		Lang:     locale.Code,
		DarkMode: ctx.AS.DarkMode(c),
		Title:    ctx.TS.T(locale.Code, "page.register.title"),
		Heading:  ctx.TS.T(locale.Code, "page.register.header"),
		Alerts:   ctx.AS.Alerts(c),
		Locale:   locale,
		Locales:  locales,
		Form: form.View{
			Action: "/register",
			Method: "post",
			Fields: []field.View{
				{
					Label:      ctx.TS.T(locale.Code, "page.register.fields.email"),
					Type:       "email",
					Name:       "email",
					Value:      "",
					Errors:     []string{},
					FieldClass: "admin-login__field",
					InputClass: "admin-login__field__input",
				},
				{
					Label:      ctx.TS.T(locale.Code, "page.register.fields.password"),
					Type:       "password",
					Name:       "password",
					Value:      "",
					Errors:     []string{},
					FieldClass: "admin-login__field",
					InputClass: "admin-login__field__input",
				},
				{
					Label:      ctx.TS.T(locale.Code, "page.register.fields.confirm_password"),
					Type:       "password",
					Name:       "confirm_password",
					Value:      "",
					Errors:     []string{},
					FieldClass: "admin-login__field",
					InputClass: "admin-login__field__input",
				},
			},
			Submit: &btn.View{
				Text: ctx.TS.T(locale.Code, "page.register.submit"),
			},
			ResetPassword: &btn.View{
				Href: "/reset-password",
				Text: ctx.TS.T(locale.Code, "page.register.reset_password"),
			},
			Login: &btn.View{
				Href: "/login",
				Text: ctx.TS.T(locale.Code, "page.register.login"),
			},
			Errors: []string{},
			//Text: "",
		},
		Back: &btn.View{Text: ctx.TS.T(locale.Code, "page.register.back"), Href: "/login"},
	}, nil
}
