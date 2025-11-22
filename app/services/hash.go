package services

type HashService struct {
}

func (s *HashService) VerifyPassword(password, hash string) (bool, error) {
	return false, nil
}

func (s *HashService) HashPassword(password string) (string, error) {
	return "", nil
}
