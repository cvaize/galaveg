package users

import "galaveg/internal/modules/errors"

type Service struct {
}

func NewService() (*Service, error) {
	return &Service{}, nil
}

func (s *Service) FirstByEmail(email string) (*User, *errors.Error) {
	return nil, nil
}

func (s *Service) ExistsByEmail(email string) (bool, *errors.Error) {
	return false, nil
}

func (s *Service) Create(data *User) (UserID, *errors.Error) {
	return 0, nil
}

func (s *Service) Update(data *User, fields []string) (UserID, *errors.Error) {
	return 0, nil
}
