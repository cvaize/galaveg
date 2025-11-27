package services

import "galaveg/config"

type HashService struct {
	cfg *config.Config
}

func NewHashService(cfg *config.Config) (*HashService, error) {
	return &HashService{cfg}, nil
}

func (s *HashService) VerifyPassword(password, hash string) (bool, error) {
	return false, nil
}

func (s *HashService) HashPassword(password string) (string, error) {
	return "", nil
}
