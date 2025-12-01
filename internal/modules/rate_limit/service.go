package rate_limit

import (
	"galaveg/internal/modules/errors"
	"galaveg/internal/modules/translator"
	"time"
)

type Service = *ServiceImpl

type ServiceImpl struct {
	ts translator.Service
}

func NewService(ts translator.Service) (*ServiceImpl, *errors.Error) {
	return &ServiceImpl{ts}, nil
}

func (s *ServiceImpl) Attempt(key string, maxAttempts int, ttl time.Duration) (bool, *errors.Error) {
	return true, nil
}

func (s *ServiceImpl) TtlMessage(locale, key string) (string, *errors.Error) {
	return "", nil
}

func (s *ServiceImpl) Clear(key string) *errors.Error {
	return nil
}
