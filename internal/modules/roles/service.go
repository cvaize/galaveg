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

func (r *ServiceImpl) Paginate(page int, perPage int, filterValues []interface{}, whereClauses []string, columns []string, orderBy string) ([]*RoleDto, int64, int, error) {
	return r.dbRepo.Paginate(page, perPage, filterValues, whereClauses, columns, orderBy)
}
