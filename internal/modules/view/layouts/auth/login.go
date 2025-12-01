package auth

import (
	"galaveg/internal/modules/alerts"
	"galaveg/internal/modules/app"
	localesModule "galaveg/internal/modules/locales"
	"galaveg/internal/modules/translator"
	"galaveg/internal/modules/view/components/btn"
	"galaveg/internal/modules/view/components/field"
	"galaveg/internal/modules/view/layouts/auth/form"
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

func NewLogin(c *gin.Context, as app.Service, ls *localesModule.ServiceImpl, ts translator.Service, s sessions.Session, data *LoginViewData) (*View, error) {
	locale := ls.GetLocale(ls.Locale(c, nil))
	locales := ls.GetLocales()

	return &View{
		Lang:     locale.Code,
		DarkMode: as.DarkMode(c),
		Title:    ts.T(locale.Code, "page.login.title"),
		Heading:  ts.T(locale.Code, "page.login.header"),
		Alerts:   alerts.Flashes(s),
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
					Label:      ts.T(locale.Code, "page.login.fields.email"),
					Type:       "email",
					Name:       "email",
					Value:      data.EmailValue,
					Errors:     data.EmailErrors,
					FieldClass: "admin-login__field",
					InputClass: "admin-login__field__input",
				},
				{
					Label:      ts.T(locale.Code, "page.login.fields.password"),
					Type:       "password",
					Name:       "password",
					Value:      data.PasswordValue,
					Errors:     data.PasswordErrors,
					FieldClass: "admin-login__field",
					InputClass: "admin-login__field__input",
				},
			},
			Submit: &btn.View{
				Text: ts.T(locale.Code, "page.login.submit"),
			},
			ResetPassword: &btn.View{
				Href: "/reset-password",
				Text: ts.T(locale.Code, "page.login.reset_password"),
			},
			Register: &btn.View{
				Href: "/register",
				Text: ts.T(locale.Code, "page.login.register"),
			},
			Errors: data.Errors,
		},
	}, nil
}
