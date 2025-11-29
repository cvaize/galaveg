package context

import (
	"galaveg/config"
	"galaveg/internal/infrastructures/db"
	"galaveg/internal/infrastructures/html"
	"galaveg/internal/infrastructures/mail"
	"galaveg/internal/infrastructures/session"
	"galaveg/internal/modules/app"
	"galaveg/internal/modules/hash"
	"galaveg/internal/modules/locales"
	mailModule "galaveg/internal/modules/mail"
	"galaveg/internal/modules/notifications"
	"galaveg/internal/modules/roles"
	"galaveg/internal/modules/template"
	"galaveg/internal/modules/translator"
	"galaveg/internal/modules/users"
	"galaveg/pkg/utils"
)

type Context struct {
	Cfg   *config.Config
	Infra struct {
		//Cache *sql.DB
		Html         html.HtmlEngine
		Db           db.Db
		Mail         mail.Mail
		SessionStore session.SessionStore
	}
	//Repositories struct{}
	Services struct {
		App           *app.Service
		Users         *users.Service
		Translator    *translator.Service
		Template      *template.Service
		Roles         *roles.Service
		Mail          *mailModule.Service
		Notifications *notifications.Service
		Locales       *locales.Service
		Hash          *hash.Service
	}
}

//goland:noinspection GoUnhandledErrorResult
func (ctx *Context) Close() {
	db.Close(ctx.Infra.Db)
	mail.Close(ctx.Infra.Mail)
}

func Must(cfg *config.Config) *Context {
	ctx := &Context{}
	ctx.Cfg = cfg
	ctx.Infra.Db = utils.Must(db.New(cfg))
	ctx.Infra.Html = utils.Must(html.New(cfg))
	ctx.Infra.Mail = utils.Must(mail.New(cfg))
	ctx.Infra.SessionStore = utils.Must(session.NewStore(cfg))

	ctx.Services.App = utils.Must(app.NewService(cfg))
	ctx.Services.Users = utils.Must(users.NewService())
	// TODO: Добавить в конфигурацию "resources/translates/" и локали
	ctx.Services.Translator = utils.Must(translator.NewServiceFromFiles(cfg.GetFolder("resources/translates/"), cfg.App.Locale))
	ctx.Services.Template = utils.Must(template.NewService(ctx.Infra.Html))
	ctx.Services.Roles = utils.Must(roles.NewService())
	ctx.Services.Mail = utils.Must(mailModule.NewService(ctx.Infra.Mail))
	ctx.Services.Notifications = utils.Must(notifications.NewService(ctx.Services.Mail))
	ctx.Services.Locales = utils.Must(locales.NewService(cfg, []locales.Locale{
		{Code: "en", ShortName: "en", FullName: "English"},
		{Code: "ru", ShortName: "ru", FullName: "Русский"},
	}))
	ctx.Services.Hash = utils.Must(hash.NewService())
	return ctx
}
