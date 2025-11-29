package roles

import "galaveg/internal/modules/errors"

type Service struct {
}

func NewService() (*Service, *errors.Error) {
	return &Service{}, nil
}
