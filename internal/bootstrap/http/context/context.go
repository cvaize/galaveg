package context

import (
	"galaveg/config"
	"galaveg/internal/infrastructures/db"
	"galaveg/internal/infrastructures/html"
	"galaveg/internal/infrastructures/mail"
	"galaveg/internal/infrastructures/session"
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
	//Services struct {
	//	Users    *services.UserService
	//}
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

	//ctx.S.ES = utils.Must(services.NewErrorService())
	//ctx.S.AlS = utils.Must(services.NewAlertService(ctx.S.ES))
	//ctx.S.LS = utils.Must(services.NewLocaleService([]dto.Locale{
	//	{Code: "en", ShortName: "en", FullName: "English"},
	//	{Code: "ru", ShortName: "ru", FullName: "Русский"},
	//}))
	//ctx.S.AS = utils.Must(services.NewAppService(cfg, ctx.S.LS))
	//ctx.S.US = utils.Must(services.NewUserService())
	//ctx.S.HS = utils.Must(services.NewHashService())
	//ctx.S.AuthS = utils.Must(services.NewAuthService(cfg, ctx.S.AS, ctx.S.US, ctx.S.HS, ctx.S.ES))
	//ctx.S.MS = utils.Must(services.NewMailService())
	//ctx.S.RS = utils.Must(services.NewRoleService())
	//ctx.S.SS = utils.Must(services.NewSessionService(cfg, ctx.S.ES))
	//ctx.S.TplS = utils.Must(services.NewTemplateService(cfg, ctx.Infra.Html))
	//// TODO: Добавить в конфигурацию "resources/translates/" и локали
	//ctx.S.TS = utils.Must(services.NewTranslatorServiceFromFiles(cfg.GetFolder("resources/translates/"), cfg.App.Locale, ctx.S.ES))
	return ctx
}
