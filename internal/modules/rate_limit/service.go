package rate_limit

import (
	"galaveg/internal/modules/errors"
	"galaveg/internal/modules/translator"
	"time"
)

type Service struct {
	ts *translator.Service
}

func NewService(ts *translator.Service) (Service, *errors.Error) {
	return Service{ts}, nil
}

func (s *Service) Attempt(key string, maxAttempts int, ttl time.Duration) (bool, *errors.Error) {
	return true, nil
}

func (s *Service) TtlMessage(locale, key string) (string, *errors.Error) {
	return "", nil
}

func (s *Service) Clear(key string) *errors.Error {
	return nil
}
