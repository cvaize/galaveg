package roles

import (
	"galaveg/internal/modules/errors"
)

type Service = *ServiceImpl

type ServiceImpl struct {
	dbRepo DbRepo
}

func NewService(dbRepo DbRepo) (*ServiceImpl, *errors.Error) {
	return &ServiceImpl{dbRepo}, nil
}

func (r *ServiceImpl) First(values []interface{}, whereClauses []string, columns []string, orderBy string) (*RoleDto, error) {
	return r.dbRepo.First(values, whereClauses, columns, orderBy)
}
