package singleton

import (
	"database/sql"
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

func init() {
	C = config.New(filepath.Join(viper.GetString("APP_FOLDER"), ".env"))
	DB = db.New(C.Db)
	TS = services.MustTranslatorServiceFromFiles(C.GetFolder("resources/lang/"), C.App.Locale)
}
