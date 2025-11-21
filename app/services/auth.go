package services

import "galaveg/app/dto"

type AuthService struct {
}

func (s AuthService) Login(email, password string) (dto.UserID, error) {

	return 0, nil
}

func (s AuthService) Register(email, password string) (dto.UserID, error) {

	return 0, nil
}
