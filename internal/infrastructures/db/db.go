package db

import (
	"database/sql"
	"fmt"
	"galaveg/config"
	"galaveg/pkg/logger"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Db = *sql.DB

func New(cfg *config.Config) (Db, error) {
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

func Close(db Db) error {
	err := db.Close()
	if err != nil {
		logger.Errorf("Db Close() error: %s", err)
	}

	return err
}
