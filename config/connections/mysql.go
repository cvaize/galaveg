package connections

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

func init() {
	viper.SetDefault("MYSQL_HOST", "localhost")
	viper.SetDefault("MYSQL_PORT", 3306)
	viper.SetDefault("MYSQL_DATABASE", "database")
	viper.SetDefault("MYSQL_USERNAME", "database_user")
	viper.SetDefault("MYSQL_PASSWORD", "database_password")
	// MYSQL_PREFIX = glg_
	viper.SetDefault("MYSQL_PREFIX", "")
	// MYSQL_TLS = true, false, skip-verify, preferred, <name>
	viper.SetDefault("MYSQL_TLS", "false")
	// MYSQL_SOCKET = /var/run/mysqld/mysqld.sock
	viper.SetDefault("MYSQL_SOCKET", "")
	// MYSQL_CHARSET = utf8mb4
	viper.SetDefault("MYSQL_CHARSET", "utf8mb4")

	viper.SetDefault("MYSQL_MAX_OPEN_CONNS", 25)
	viper.SetDefault("MYSQL_MAX_IDLE_CONNS", 25)
	viper.SetDefault("MYSQL_CONN_MAX_LIFETIME", int64(time.Hour))
	viper.SetDefault("MYSQL_CONN_MAX_IDLE_TIME", int64(5*time.Minute))
}

type MysqlConfig struct {
	Host            string
	Port            int
	Database        string
	Username        string
	Password        string
	Prefix          string
	Tls             string
	Socket          string
	Charset         string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime int64
	ConnMaxIdleTime int64
}

func (c *MysqlConfig) DSN() string {
	protocol := "tcp"
	address := fmt.Sprintf("%s:%d", c.Host, c.Port)

	if c.Socket != "" {
		protocol = "unix"
		address = c.Socket
	}

	return fmt.Sprintf(`%s:%s@%s(%s)/%s?charset=%s&tls=%s&multiStatements=true`, c.Username, c.Password, protocol, address, c.Database, c.Charset, c.Tls)
}

func NewMysqlConfig() *MysqlConfig {
	return &MysqlConfig{
		Host:            viper.GetString("MYSQL_HOST"),
		Port:            viper.GetInt("MYSQL_PORT"),
		Database:        viper.GetString("MYSQL_DATABASE"),
		Username:        viper.GetString("MYSQL_USERNAME"),
		Password:        viper.GetString("MYSQL_PASSWORD"),
		Prefix:          viper.GetString("MYSQL_PREFIX"),
		Tls:             viper.GetString("MYSQL_TLS"),
		Socket:          viper.GetString("MYSQL_SOCKET"),
		Charset:         viper.GetString("MYSQL_CHARSET"),
		MaxOpenConns:    viper.GetInt("MYSQL_MAX_OPEN_CONNS"),
		MaxIdleConns:    viper.GetInt("MYSQL_MAX_IDLE_CONNS"),
		ConnMaxLifetime: viper.GetInt64("MYSQL_CONN_MAX_LIFETIME"),
		ConnMaxIdleTime: viper.GetInt64("MYSQL_CONN_MAX_IDLE_TIME"),
	}
}
