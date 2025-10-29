package connections

import (
	"database/sql"
	"galaveg/config"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func InitMySQL() *sql.DB {
	dsn := config.Config.Mysql.DSN()

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// Проверка соединения
	if err := db.Ping(); err != nil {
		panic(err)
	}

	// Настройка пула
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(time.Hour)
	db.SetConnMaxIdleTime(5 * time.Minute)

	return db
}
