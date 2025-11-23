package providers

import (
	"database/sql"
	"galaveg/app/dto"
	"galaveg/app/services"
	"galaveg/config"
	"galaveg/utils/logger"
)

type Context struct {
	Cfg   *config.Config
	Db    *sql.DB
	Cache *sql.DB
	TS    *services.TranslatorService
	LS    *services.LocaleService
	AS    *services.AppService
	Auth  *services.AuthService
	RS    *services.RoleService
	ES    *services.ErrorService
	SS    *services.SessionService
	AlS   *services.AlertService
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

func NewContext(cfg *config.Config) (*Context, error) {
	DB, err := NewDB(cfg)
	if err != nil {
		return nil, err
	}

	// TODO: Добавить в конфигурацию "resources/translates/" и локали
	TS := services.MustTranslatorServiceFromFiles(cfg.GetFolder("resources/translates/"), cfg.App.Locale)
	LS := services.MustLocaleService([]dto.Locale{
		{Code: "en", ShortName: "en", FullName: "English"},
		{Code: "ru", ShortName: "ru", FullName: "Русский"},
	})
	RS := services.MustRoleService()
	AS := services.MustAppService(cfg, LS, RS, TS)
	return &Context{
		Cfg: cfg,
		Db:  DB,
		TS:  TS,
		LS:  LS,
		RS:  RS,
		AS:  AS,
	}, nil
}

func MustContext(C *config.Config) *Context {
	c, e := NewContext(C)
	if e != nil {
		panic(e)
	}
	return c
}
