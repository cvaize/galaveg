package providers

import (
	"database/sql"
	"fmt"
	"galaveg/app/connections/mysql"
	"galaveg/app/dto"
	"galaveg/app/services"
	"galaveg/config"
	"galaveg/utils/logger"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
)

type Context struct {
	C     *config.Config
	DB    *sql.DB
	Cache *sql.DB
	TS    *services.TranslatorService
	LS    *services.LocaleService
	AS    *services.AppService
	RS    *services.RoleService
}

func (ctx *Context) Close() {
	if ctx.DB != nil {
		err := ctx.DB.Close()
		logger.Errorf("DB Close() error: %s", err)
	}
	if ctx.Cache != nil {
		err := ctx.Cache.Close()
		logger.Errorf("Cache Close() error: %s", err)
	}
}

func NewContext(C *config.Config) (*Context, error) {
	DB, err := mysql.New(C.Connections.Mysql)
	if err != nil {
		return nil, err
	}

	// TODO: Добавить в конфигурацию "resources/translates/" и локали
	TS := services.MustTranslatorServiceFromFiles(C.GetFolder("resources/translates/"), C.App.Locale)
	LS := services.MustLocaleService([]dto.Locale{
		{Code: "en", ShortName: "en", FullName: "English"},
		{Code: "ru", ShortName: "ru", FullName: "Русский"},
	})
	RS := services.MustRoleService()
	AS := services.MustAppService(C.App, LS, RS, TS)
	return &Context{
		C:  C,
		DB: DB,
		TS: TS,
		LS: LS,
		RS: RS,
		AS: AS,
	}, nil
}

func MustContext(C *config.Config) *Context {
	c, e := NewContext(C)
	if e != nil {
		panic(e)
	}
	return c
}

func (ctx *Context) NewSessionStore() (sessions.Store, error) {
	if ctx.C.Session.Store == "redis" && ctx.C.Session.StoreConnection == "redis" {
		c := ctx.C.Connections.Redis
		// TODO: Добавить в конфиг настроек отдельно ctx.C.App.Key и настройки подключения к базе данных
		return redis.NewStore(100, "tcp", c.Host, c.Username, c.Password, []byte(ctx.C.App.Key))
	}

	return nil, fmt.Errorf("session store settings should have the values: \"redis\"")
}
