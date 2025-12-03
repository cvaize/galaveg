package roles

import (
	dbModule "galaveg/internal/modules/db"
	"galaveg/internal/modules/errors"
)

type Service = *ServiceImpl

type ServiceImpl struct {
	dbRepo DbRepo
}

func NewService(dbRepo DbRepo) (*ServiceImpl, *errors.Error) {
	return &ServiceImpl{dbRepo}, nil
}

func (r *ServiceImpl) All(query *dbModule.DbRepoQuery) ([]*RoleDto, error) {
	return r.dbRepo.All(query)
}
