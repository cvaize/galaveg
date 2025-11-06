package singleton

import (
	"database/sql"
	"galaveg/config"
	"galaveg/connections/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"path/filepath"
)

var c config.Config
var C *config.Config
var DB *sql.DB

func init() {
	c = config.New(filepath.Join(viper.GetString("APP_FOLDER"), ".env"))
	C = &c
	DB = db.New(C.Db)
}
