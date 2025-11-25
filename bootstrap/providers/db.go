package providers

import (
	"database/sql"
	"fmt"
	"galaveg/config"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func NewDB(cfg *config.Config) (*sql.DB, error) {
	c := cfg.Db
	protocol := "tcp"
	address := fmt.Sprintf("%s:%d", c.Host, c.Port)

	if c.Socket != "" {
		protocol = "unix"
		address = c.Socket
	}

	dsn := fmt.Sprintf(`%s:%s@%s(%s)/%s?charset=%s&tls=%s&multiStatements=true`, c.Username, c.Password, protocol, address, c.Database, c.Charset, c.Tls)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Настройка пула
	db.SetMaxOpenConns(c.MaxOpenConns)
	db.SetMaxIdleConns(c.MaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(c.ConnMaxLifetime))
	db.SetConnMaxIdleTime(time.Duration(c.ConnMaxIdleTime))

	return db, nil
}

func MustDB(cfg *config.Config) *sql.DB {
	db, e := NewDB(cfg)
	if e != nil {
		panic(e)
	}
	return db
}
