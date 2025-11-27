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
	Cfg   *config.Config
	Html  *template.Template
	Db    *sql.DB
	Cache *sql.DB
	Mail  *mail.Client
	TS    *services.TranslatorService
	LS    *services.LocaleService
	AS    *services.AppService
	Auth  *services.AuthService
	RS    *services.RoleService
	ES    *services.ErrorService
	SS    *services.SessionService
	AlS   *services.AlertService
	MS    *services.MailService
	US    *services.UserService
	HS    *services.HashService
	TplS  *services.TemplateService
}

func (ctx *Context) Close() {
	if ctx.Db != nil {
		err := ctx.Db.Close()
		logger.Errorf("Db Close() error: %s", err)
	}
	if ctx.Cache != nil {
		err := ctx.Cache.Close()
		logger.Errorf("Cache Close() error: %s", err)
	}
}

func MustContext(cfg *config.Config) *Context {
	html := utils.Must(NewHtmlEngine(cfg))
	db := utils.Must(NewDB(cfg))
	mc := utils.Must(NewMail(cfg))

	es := utils.Must(services.NewErrorService())
	tpl := utils.Must(services.NewTemplateService(cfg, html))
	// TODO: Добавить в конфигурацию "resources/translates/" и локали
	ts := utils.Must(services.NewTranslatorServiceFromFiles(cfg.GetFolder("resources/translates/"), cfg.App.Locale))
	ls := utils.Must(services.NewLocaleService([]dto.Locale{
		{Code: "en", ShortName: "en", FullName: "English"},
		{Code: "ru", ShortName: "ru", FullName: "Русский"},
	}))
	rs := utils.Must(services.NewRoleService())
	as := utils.Must(services.NewAppService(cfg, ls, rs, ts))
	ms := utils.Must(services.NewMailService(cfg, mc))
	us := utils.Must(services.NewUserService(cfg))
	hs := utils.Must(services.NewHashService(cfg))
	ss := utils.Must(services.NewSessionService(cfg, es))
	als := utils.Must(services.NewAlertService(es))
	auth := utils.Must(services.NewAuthService(cfg, us, ts, hs, es))
	return &Context{
		Cfg:  cfg,
		Html: html,
		Db:   db,
		Mail: mc,
		TS:   ts,
		LS:   ls,
		AS:   as,
		Auth: auth,
		RS:   rs,
		MS:   ms,
		US:   us,
		HS:   hs,
		ES:   es,
		SS:   ss,
		AlS:  als,
		TplS: tpl,
	}
}
