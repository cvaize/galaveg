package hash

import "galaveg/internal/modules/errors"

//import "golang.org/x/crypto/argon2"

type Service = *ServiceImpl

type ServiceImpl struct {
}

func NewService() (*ServiceImpl, *errors.Error) {
	return &ServiceImpl{}, nil
}

func (s *ServiceImpl) VerifyPassword(password, hash string) (bool, *errors.Error) {
	return false, nil
}

func (s *ServiceImpl) HashPassword(password string) (string, *errors.Error) {
	//argon2.IDKey()
	return "", nil
}
