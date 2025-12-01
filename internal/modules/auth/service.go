package auth

import (
	"galaveg/internal/modules/errors"
	"galaveg/internal/modules/hash"
	"galaveg/internal/modules/users"
	"galaveg/pkg/debug"
)

type Service struct {
	hs *hash.Service
}

func NewService(hs *hash.Service) (Service, *errors.Error) {
	return Service{hs}, nil
}

func (s *Service) UpdatePassword(id users.UserID, password string) *errors.Error {
	passwordHashed, e := s.hs.HashPassword(password)
	if e != nil {
		return errors.E500(e, "auth.Service.UpdatePassword.HashPassword", "")
	}

	debug.Dump(passwordHashed)

	// TODO: users.Repo.updateById set passwordHashed

	return nil
}
