package singleton

import (
	"database/sql"
	"galaveg/config"
	"galaveg/connections/db"
	_ "github.com/go-sql-driver/mysql"
)

var c config.Config
var C *config.Config
var DB *sql.DB

func init() {
	c = config.New()
	C = &c
	DB = db.New(C.Db)
}
