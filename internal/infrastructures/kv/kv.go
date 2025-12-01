package kv

import (
	"fmt"
	"galaveg/internal/config"
	"galaveg/pkg/logger"
	"github.com/redis/go-redis/v9"
)

type KV = *redis.Client

func New(cfg *config.Config) (KV, error) {
	addr := fmt.Sprintf("%s:%d", cfg.KeyValue.RedisHost, cfg.KeyValue.RedisPort)
	username := cfg.KeyValue.RedisUsername
	password := cfg.KeyValue.RedisPassword

	kv := redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: username,
		Password: password,
		DB:       0,
	})

	return kv, nil
}

func Close(db KV) error {
	err := db.Close()
	if err != nil {
		logger.Errorf("Db Close() error: %s", err)
	}

	return err
}
