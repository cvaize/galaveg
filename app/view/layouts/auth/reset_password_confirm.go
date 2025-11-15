package auth

import (
	"galaveg/app/view/components/btn"
	"galaveg/app/view/components/field"
	"galaveg/app/view/layouts/auth/form"
	"galaveg/bootstrap/singleton"
	"github.com/gin-gonic/gin"
)

func NewResetPasswordConfirm(c *gin.Context) (*View, error) {
	locale := singleton.LS.GetLocale(singleton.AS.Locale(c, nil))
	locales := singleton.LS.GetLocales()

	return &View{
		Lang:     locale.Code,
		DarkMode: singleton.AS.DarkMode(c),
		Title:    singleton.TS.T(locale.Code, "page.reset_password_confirm.title"),
		Heading:  singleton.TS.T(locale.Code, "page.reset_password_confirm.header"),
		Alerts:   singleton.AS.Alerts(c),
		Locale:   locale,
		Locales:  locales,
		Form: form.View{
			///reset-password-confirm?code={}&email={}
			Action: "/reset-password-confirm",
			Method: "post",
			Fields: []field.View{
				{
					Type:       "hidden",
					Name:       "code",
					Value:      "",
					FieldClass: "admin-login__field",
					InputClass: "admin-login__field__input",
				},
				{
					Label:      singleton.TS.T(locale.Code, "page.reset_password_confirm.fields.email"),
					Type:       "email",
					Name:       "email",
					Value:      "",
					Errors:     []string{},
					Readonly:   true,
					FieldClass: "admin-login__field",
					InputClass: "admin-login__field__input",
				},
				{
					Label:      singleton.TS.T(locale.Code, "page.reset_password_confirm.fields.password"),
					Type:       "password",
					Name:       "password",
					Value:      "",
					Errors:     []string{},
					FieldClass: "admin-login__field",
					InputClass: "admin-login__field__input",
				},
				{
					Label:      singleton.TS.T(locale.Code, "page.reset_password_confirm.fields.confirm_password"),
					Type:       "password",
					Name:       "confirm_password",
					Value:      "",
					Errors:     []string{},
					FieldClass: "admin-login__field",
					InputClass: "admin-login__field__input",
				},
			},
			Submit: &btn.View{
				Text: singleton.TS.T(locale.Code, "page.reset_password_confirm.submit"),
			},
			Errors: []string{},
		},
		Back: &btn.View{Text: singleton.TS.T(locale.Code, "page.reset_password_confirm.back"), Href: "/reset-password"},
	}, nil
}
