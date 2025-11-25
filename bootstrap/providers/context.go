package providers

import (
	"database/sql"
	"galaveg/app/dto"
	"galaveg/app/services"
	"galaveg/config"
	"galaveg/utils/logger"
	"github.com/wneessen/go-mail"
)

type Context struct {
	Cfg   *config.Config
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
	db := MustDB(cfg)
	mc := MustMail(cfg)

	es := services.MustErrorService()
	// TODO: Добавить в конфигурацию "resources/translates/" и локали
	ts := services.MustTranslatorServiceFromFiles(cfg.GetFolder("resources/translates/"), cfg.App.Locale)
	ls := services.MustLocaleService([]dto.Locale{
		{Code: "en", ShortName: "en", FullName: "English"},
		{Code: "ru", ShortName: "ru", FullName: "Русский"},
	})
	rs := services.MustRoleService()
	as := services.MustAppService(cfg, ls, rs, ts)
	ms := services.MustMailService(cfg, mc)
	us := services.MustUserService(cfg)
	hs := services.MustHashService(cfg)
	ss := services.MustSessionService(cfg, es)
	als := services.MustAlertService(es)
	auth := services.MustAuthService(cfg, us, ts, hs, es)
	return &Context{
		Cfg:  cfg,
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
	}
}
