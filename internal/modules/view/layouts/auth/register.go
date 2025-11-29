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

type RegisterViewData struct {
	EmailValue            string
	EmailErrors           []string
	PasswordValue         string
	PasswordErrors        []string
	ConfirmPasswordValue  string
	ConfirmPasswordErrors []string
	Errors                []string
}

func NewRegister(c *gin.Context, as *app.Service, ls *localesModule.Service, ts *translator.Service, s sessions.Session, data *RegisterViewData) (*View, error) {
	locale := ls.GetLocale(ls.Locale(c, nil))
	locales := ls.GetLocales()

	return &View{
		Lang:     locale.Code,
		DarkMode: as.DarkMode(c),
		Title:    ts.T(locale.Code, "page.register.title"),
		Heading:  ts.T(locale.Code, "page.register.header"),
		Alerts:   alerts.Flashes(s),
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
					Label:      ts.T(locale.Code, "page.register.fields.email"),
					Type:       "email",
					Name:       "email",
					Value:      data.EmailValue,
					Errors:     data.EmailErrors,
					FieldClass: "admin-login__field",
					InputClass: "admin-login__field__input",
				},
				{
					Label:      ts.T(locale.Code, "page.register.fields.password"),
					Type:       "password",
					Name:       "password",
					Value:      data.PasswordValue,
					Errors:     data.PasswordErrors,
					FieldClass: "admin-login__field",
					InputClass: "admin-login__field__input",
				},
				{
					Label:      ts.T(locale.Code, "page.register.fields.confirm_password"),
					Type:       "password",
					Name:       "confirm_password",
					Value:      data.ConfirmPasswordValue,
					Errors:     data.ConfirmPasswordErrors,
					FieldClass: "admin-login__field",
					InputClass: "admin-login__field__input",
				},
			},
			Submit: &btn.View{
				Text: ts.T(locale.Code, "page.register.submit"),
			},
			ResetPassword: &btn.View{
				Href: "/reset-password",
				Text: ts.T(locale.Code, "page.register.reset_password"),
			},
			Login: &btn.View{
				Href: "/login",
				Text: ts.T(locale.Code, "page.register.login"),
			},
			Errors: data.Errors,
		},
		Back: &btn.View{Text: ts.T(locale.Code, "page.register.back"), Href: "/login"},
	}, nil
}
