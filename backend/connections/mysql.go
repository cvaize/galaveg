package connections

import (
	"database/sql"
	"galaveg/config"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var MySQL *sql.DB

func init() {
	dsn := config.Config.Mysql.DSN()

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// Настройка пула
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(time.Hour)
	db.SetConnMaxIdleTime(5 * time.Minute)

	MySQL = db
}
