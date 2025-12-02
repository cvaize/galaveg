package users

import "galaveg/internal/modules/errors"

type Service = *ServiceImpl

type ServiceImpl struct {
}

func NewService() (*ServiceImpl, *errors.Error) {
	return &ServiceImpl{}, nil
}

func (s *ServiceImpl) FirstByEmail(email string) (*UserDto, *errors.Error) {
	return nil, nil
}

func (s *ServiceImpl) ExistsByEmail(email string) (bool, *errors.Error) {
	return false, nil
}

func (s *ServiceImpl) Create(data *UserDto) (ID, *errors.Error) {
	// TODO: Email to LowCase
	return 0, nil
}

func (s *ServiceImpl) Update(data *UserDto, fields []string) (ID, *errors.Error) {
	// TODO: Email to LowCase
	return 0, nil
}
