package providers

import (
	"database/sql"
	"galaveg/app/dto"
	"galaveg/app/services"
	"galaveg/config"
	"galaveg/utils"
	"galaveg/utils/logger"
	"github.com/wneessen/go-mail"
	"html/template"
)

type Context struct {
	Cfg *config.Config
	// Infrastructures
	Infra struct {
		Html *template.Template
		Db   *sql.DB
		//Cache *sql.DB
		Mail *mail.Client
	}
	// Repositories
	R struct{}
	// Services
	S struct {
		AlS   *services.AlertService
		AS    *services.AppService
		AuthS *services.AuthService
		ES    *services.ErrorService
		HS    *services.HashService
		LS    *services.LocaleService
		MS    *services.MailService
		RS    *services.RoleService
		SS    *services.SessionService
		TplS  *services.TemplateService
		TS    *services.TranslatorService
		US    *services.UserService
	}
}

func (ctx *Context) Close() {
	if ctx.Infra.Db != nil {
		err := ctx.Infra.Db.Close()
		logger.Errorf("Db Close() error: %s", err)
	}
	//if ctx.Infra.Cache != nil {
	//	err := ctx.Infra.Cache.Close()
	//	logger.Errorf("Cache Close() error: %s", err)
	//}
	if ctx.Infra.Mail != nil {
		err := ctx.Infra.Mail.Close()
		logger.Errorf("Mail Close() error: %s", err)
	}
}

func MustContext(cfg *config.Config) *Context {
	ctx := &Context{}
	ctx.Cfg = cfg
	ctx.Infra.Db = utils.Must(NewDB(cfg))
	ctx.Infra.Html = utils.Must(NewHtmlEngine(cfg))
	ctx.Infra.Mail = utils.Must(NewMail(cfg))

	ctx.S.ES = utils.Must(services.NewErrorService())
	ctx.S.AlS = utils.Must(services.NewAlertService(ctx.S.ES))
	ctx.S.LS = utils.Must(services.NewLocaleService([]dto.Locale{
		{Code: "en", ShortName: "en", FullName: "English"},
		{Code: "ru", ShortName: "ru", FullName: "Русский"},
	}))
	ctx.S.AS = utils.Must(services.NewAppService(cfg, ctx.S.LS))
	ctx.S.US = utils.Must(services.NewUserService())
	ctx.S.HS = utils.Must(services.NewHashService())
	ctx.S.AuthS = utils.Must(services.NewAuthService(cfg, ctx.S.AS, ctx.S.US, ctx.S.HS, ctx.S.ES))
	ctx.S.MS = utils.Must(services.NewMailService())
	ctx.S.RS = utils.Must(services.NewRoleService())
	ctx.S.SS = utils.Must(services.NewSessionService(cfg, ctx.S.ES))
	ctx.S.TplS = utils.Must(services.NewTemplateService(cfg, ctx.Infra.Html))
	// TODO: Добавить в конфигурацию "resources/translates/" и локали
	ctx.S.TS = utils.Must(services.NewTranslatorServiceFromFiles(cfg.GetFolder("resources/translates/"), cfg.App.Locale, ctx.S.ES))
	return ctx
}
