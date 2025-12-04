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

func (r *ServiceImpl) Count(filterValues []interface{}, whereClauses []string) (int64, error) {
	return r.dbRepo.Count(filterValues, whereClauses)
}
