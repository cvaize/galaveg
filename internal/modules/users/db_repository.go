package users

import (
	"galaveg/internal/config"
	"galaveg/internal/infrastructures/db"
)

type DbRepo = *DbRepoImpl

type DbRepoImpl struct {
	cfg *config.Config
	db  db.Db
}

func NewDbRepoImpl(cfg *config.Config, db db.Db) *DbRepoImpl {
	return &DbRepoImpl{cfg, db}
}
