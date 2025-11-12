package singleton

import (
	"database/sql"
	"galaveg/app/dto"
	"galaveg/app/services"
	"galaveg/config"
	"galaveg/connections/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"path/filepath"
)

var C *config.Config
var DB *sql.DB
var TS *services.TranslatorService
var LS *services.LocaleService
var AS *services.AppService
var RS *services.RoleService

func init() {
	C = config.New(filepath.Join(viper.GetString("APP_FOLDER"), ".env"))
	DB = db.New(C.Db)
	TS = services.MustTranslatorServiceFromFiles(C.GetFolder("resources/translates/"), C.App.Locale)
	LS = services.MustLocaleService([]dto.Locale{
		{Code: "en", ShortName: "en", FullName: "English"},
		{Code: "ru", ShortName: "ru", FullName: "Русский"},
	})
	RS = services.MustRoleService()
	AS = services.MustAppService(C.App, LS, RS, TS)
}
