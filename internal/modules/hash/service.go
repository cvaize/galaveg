package hash

import "galaveg/internal/modules/errors"

type Service struct {
}

func NewService() (*Service, error) {
	return &Service{}, nil
}

func (s *Service) VerifyPassword(password, hash string) (bool, *errors.Error) {
	return false, nil
}

func (s *Service) HashPassword(password string) (string, *errors.Error) {
	return "", nil
}
