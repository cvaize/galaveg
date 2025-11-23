package services

import "galaveg/app/dto"

type UserService struct {
}

func (s *UserService) FirstByEmail(email string) (*dto.User, error) {
	return nil, nil
}

func (s *UserService) ExistsByEmail(email string) (bool, error) {
	return false, nil
}

func (s *UserService) Create(data *dto.User) (dto.UserID, error) {
	return 0, nil
}

func (s *UserService) Update(data *dto.User, fields []string) (dto.UserID, error) {
	return 0, nil
}
