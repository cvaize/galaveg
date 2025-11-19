package mysql

import (
	"database/sql"
	"galaveg/config/connections"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func New(c *connections.MysqlConfig) (*sql.DB, error) {
	dsn := c.DSN()

	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Настройка пула
	DB.SetMaxOpenConns(c.MaxOpenConns)
	DB.SetMaxIdleConns(c.MaxIdleConns)
	DB.SetConnMaxLifetime(time.Duration(c.ConnMaxLifetime))
	DB.SetConnMaxIdleTime(time.Duration(c.ConnMaxIdleTime))

	return DB, nil
}

func Must(c *connections.MysqlConfig) *sql.DB {
	DB, err := New(c)
	if err != nil {
		panic(err)
	}

	return DB
}
