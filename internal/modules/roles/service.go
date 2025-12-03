package roles

import "galaveg/internal/modules/errors"

type Service = *ServiceImpl

type ServiceImpl struct {
	dbRepo DbRepo
}

func NewService(dbRepo DbRepo) (*ServiceImpl, *errors.Error) {
	return &ServiceImpl{dbRepo}, nil
}

//func (r *ServiceImpl) AllIds(filters map[string]interface{}, orderBy string) ([]ID, error) {
//	return r.dbRepo.AllIds(filters, orderBy)
//}
//
//func (r *ServiceImpl) All(filters map[string]interface{}, orderBy string, columns []string) ([]*RoleDto, error) {
//	return r.dbRepo.All(filters, orderBy, columns)
//}
//
//func (r *ServiceImpl) First(filters map[string]interface{}, columns []string) (*RoleDto, error) {
//	return r.dbRepo.First(filters, columns)
//}
