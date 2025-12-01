package roles

import "galaveg/internal/modules/errors"

type Service = *ServiceImpl

type ServiceImpl struct {
}

func NewService() (*ServiceImpl, *errors.Error) {
	return &ServiceImpl{}, nil
}
