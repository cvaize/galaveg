package auth

import (
	"galaveg/app/view/components/btn"
	"galaveg/app/view/components/field"
	"galaveg/app/view/layouts/auth/form"
	"galaveg/bootstrap/providers"
	"github.com/gin-gonic/gin"
)

func NewResetPasswordConfirm(c *gin.Context, ctx *providers.Context) (*View, error) {
	locale := ctx.LS.GetLocale(ctx.AS.Locale(c, nil))
	locales := ctx.LS.GetLocales()

	return &View{
		Lang:     locale.Code,
		DarkMode: ctx.AS.DarkMode(c),
		Title:    ctx.TS.T(locale.Code, "page.reset_password_confirm.title"),
		Heading:  ctx.TS.T(locale.Code, "page.reset_password_confirm.header"),
		Alerts:   ctx.AS.Alerts(c),
		Locale:   locale,
		Locales:  locales,
		Form: form.View{
			///reset-password-confirm?code={}&email={}
			Action: "/reset-password-confirm",
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
					Type:       "hidden",
					Name:       "code",
					Value:      "",
					FieldClass: "admin-login__field",
					InputClass: "admin-login__field__input",
				},
				{
					Label:      ctx.TS.T(locale.Code, "page.reset_password_confirm.fields.email"),
					Type:       "email",
					Name:       "email",
					Value:      "",
					Errors:     []string{},
					Readonly:   true,
					FieldClass: "admin-login__field",
					InputClass: "admin-login__field__input",
				},
				{
					Label:      ctx.TS.T(locale.Code, "page.reset_password_confirm.fields.password"),
					Type:       "password",
					Name:       "password",
					Value:      "",
					Errors:     []string{},
					FieldClass: "admin-login__field",
					InputClass: "admin-login__field__input",
				},
				{
					Label:      ctx.TS.T(locale.Code, "page.reset_password_confirm.fields.confirm_password"),
					Type:       "password",
					Name:       "confirm_password",
					Value:      "",
					Errors:     []string{},
					FieldClass: "admin-login__field",
					InputClass: "admin-login__field__input",
				},
			},
			Submit: &btn.View{
				Text: ctx.TS.T(locale.Code, "page.reset_password_confirm.submit"),
			},
			Errors: []string{},
		},
		Back: &btn.View{Text: ctx.TS.T(locale.Code, "page.reset_password_confirm.back"), Href: "/reset-password"},
	}, nil
}
