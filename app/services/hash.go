package services

type HashService struct {
}

func NewHashService() (*HashService, error) {
	return &HashService{}, nil
}

func (s *HashService) VerifyPassword(password, hash string) (bool, error) {
	return false, nil
}

func (s *HashService) HashPassword(password string) (string, error) {
	return "", nil
}
