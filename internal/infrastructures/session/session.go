package session

import (
	"fmt"
	"galaveg/internal/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
)

type Session = sessions.Session
type SessionStore = sessions.Store

func NewStore(cfg *config.Config) (SessionStore, error) {
	// TODO: Сделать команду генерации ключей приложения
	// TODO: Добавить в конфиг настроек отдельно ctx.Cfg.App.Key и настройки подключения к базе данных
	size := cfg.Session.RedisMaxIdleConnections
	url := fmt.Sprintf("%s:%d", cfg.Session.RedisHost, cfg.Session.RedisPort)
	username := cfg.Session.RedisUsername
	password := cfg.Session.RedisPassword
	key := cfg.App.Key
	return redis.NewStore(size, "tcp", url, username, password, []byte(key))
}
