package db

import (
	"database/sql"
	"galaveg/config"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func New(c config.DbConfig) *sql.DB {
	dsn := c.DSN()

	db, err := sql.Open(c.Driver, dsn)
	if err != nil {
		panic(err)
	}

	// Настройка пула
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(time.Hour)
	db.SetConnMaxIdleTime(5 * time.Minute)

	return db
}
