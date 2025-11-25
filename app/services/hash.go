package services

import "galaveg/config"

type HashService struct {
	cfg *config.Config
}

func NewHashService(cfg *config.Config) (*HashService, error) {
	return &HashService{cfg}, nil
}

func MustHashService(cfg *config.Config) *HashService {
	s, e := NewHashService(cfg)
	if e != nil {
		panic(e)
	}
	return s
}

func (s *HashService) VerifyPassword(password, hash string) (bool, error) {
	return false, nil
}

func (s *HashService) HashPassword(password string) (string, error) {
	return "", nil
}
