package users

import "galaveg/internal/modules/errors"

type Service = *ServiceImpl

type ServiceImpl struct {
	dbRepo DbRepo
}

func NewService(dbRepo DbRepo) (*ServiceImpl, *errors.Error) {
	return &ServiceImpl{dbRepo}, nil
}

func (s *ServiceImpl) FirstByEmail(email string) (*UserDto, *errors.Error) {
	return nil, nil
}

func (s *ServiceImpl) ExistsByEmail(email string) (bool, *errors.Error) {
	return false, nil
}

func (s *ServiceImpl) Create(data *UserDto) (int64, *errors.Error) {
	// TODO: Email to LowCase
	return 0, nil
}

func (s *ServiceImpl) Update(data *UserDto, fields []string) (int64, *errors.Error) {
	// TODO: Email to LowCase
	return 0, nil
}
