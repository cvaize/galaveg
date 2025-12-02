package auth

import (
	"galaveg/internal/infrastructures/db"
	"galaveg/internal/modules/errors"
)

type DbRepo = *DbRepoImpl

type DbRepoImpl struct {
	db db.Db
}

func NewDbRepoImpl(db db.Db) *DbRepoImpl {
	return &DbRepoImpl{db}
}

func (r *DbRepoImpl) FirstByEmail(email string) (*UserDto, *errors.Error) {
	return nil, nil
}

func (r *DbRepoImpl) Create(data *UserDto) *errors.Error {
	return nil
}
