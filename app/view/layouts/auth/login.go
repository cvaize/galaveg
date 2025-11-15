package auth

import (
	"galaveg/app/view/components/btn"
	"galaveg/app/view/components/field"
	"galaveg/app/view/layouts/auth/form"
	"galaveg/bootstrap/singleton"
	"github.com/gin-gonic/gin"
)

func NewLogin(c *gin.Context) (*View, error) {
	locale := singleton.LS.GetLocale(singleton.AS.Locale(c, nil))
	locales := singleton.LS.GetLocales()

	return &View{
		Lang:     locale.Code,
		DarkMode: singleton.AS.DarkMode(c),
		Title:    singleton.TS.T(locale.Code, "page.login.title"),
		Heading:  singleton.TS.T(locale.Code, "page.login.header"),
		Alerts:   singleton.AS.Alerts(c),
		Locale:   locale,
		Locales:  locales,
		Form: form.View{
			Action: "/login",
			Method: "post",
			Fields: []field.View{
				{
					Label:      singleton.TS.T(locale.Code, "page.login.fields.email"),
					Type:       "email",
					Name:       "email",
					Value:      "",
					Errors:     []string{},
					FieldClass: "admin-login__field",
					InputClass: "admin-login__field__input",
				},
				{
					Label:      singleton.TS.T(locale.Code, "page.login.fields.password"),
					Type:       "password",
					Name:       "password",
					Value:      "",
					Errors:     []string{},
					FieldClass: "admin-login__field",
					InputClass: "admin-login__field__input",
				},
			},
			Submit: &btn.View{
				Text: singleton.TS.T(locale.Code, "page.login.submit"),
			},
			ResetPassword: &btn.View{
				Href: "/reset-password",
				Text: singleton.TS.T(locale.Code, "page.login.reset_password"),
			},
			Register: &btn.View{
				Href: "/register",
				Text: singleton.TS.T(locale.Code, "page.login.register"),
			},
			Errors: []string{},
			//Text: "",
		},
		//Back:     btn.View{Text: "Назад", Href: "/login"},
	}, nil
}
