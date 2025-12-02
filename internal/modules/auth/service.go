package auth

import (
	"galaveg/internal/modules/errors"
	"galaveg/internal/modules/hash"
	"galaveg/internal/modules/users"
	"galaveg/pkg/debug"
	"strings"
)

type Service = *ServiceImpl

type ServiceImpl struct {
	hs     hash.Service
	dbRepo DbRepo
}

func NewService(hs hash.Service, dbRepo DbRepo) (*ServiceImpl, *errors.Error) {
	return &ServiceImpl{hs, dbRepo}, nil
}

func (s *ServiceImpl) UpdatePassword(id users.UserID, password string) *errors.Error {
	passwordHashed, e := s.hs.HashPassword(password)
	if e != nil {
		return errors.E500(e, "auth.ServiceImpl.UpdatePassword.HashPassword", "")
	}

	debug.Dump(passwordHashed)

	// TODO: users.Repo.updateById set passwordHashed

	return nil
}

func (s *ServiceImpl) Login(email EmailVO, password PasswordVO) (users.UserID, *errors.Error) {
	user, e := s.dbRepo.FirstByEmail(email.Value)
	if e != nil {
		// Failed to get user
		return 0, errors.E500(e, "auth.ServiceImpl.Login.FailedToGetUser", "")
	}

	if user == nil {
		// UserDto not found
		return 0, errors.E404(e, "auth.ServiceImpl.Login.UserNotFound", "")
	}

	is, e := s.hs.VerifyPassword(password.Value, user.PasswordHash)
	if e != nil {
		// Failed to verify password
		return 0, errors.E500(e, "auth.ServiceImpl.Login.FailedToVerifyPassword", "")
	}
	if !is {
		// The password does not match the saved hash
		return 0, errors.E401(e, "auth.ServiceImpl.Login.Unauthorized", "")
	}

	return user.ID, nil
}

func (s *ServiceImpl) Register(email EmailVO, password PasswordVO) *errors.Error {
	user, e := s.dbRepo.FirstByEmail(email.Value)
	if e != nil {
		// Failed to check for user existence
		return errors.E500(e, "auth.ServiceImpl.Register.FailedToGetUser", "")
	}
	if user != nil {
		if user.Email.Value != email.Value {
			return errors.E500(e, "auth.ServiceImpl.Register.EmailParamNotEqualFindEmail", "")
		}
		// This user is already registered
		return errors.E400(e, "auth.ServiceImpl.Register.DuplicateUser", "")
	}

	passwordHash, e := s.hs.HashPassword(password.Value)
	if e != nil {
		// Failed to generate password hash
		return errors.E500(e, "auth.ServiceImpl.Register.HashPasswordFail", "")
	}

	err := s.dbRepo.Create(&UserDto{Email: email, PasswordHash: passwordHash})
	if err != nil {
		eStr := err.Error()
		if strings.Contains(eStr, "Duplicate entry") {
			if strings.Contains(eStr, ".email'") {
				// There is already such a user
				return errors.E400(e, "auth.ServiceImpl.Register.DuplicateUser", "")
			}
			// There is already such a user, not intended behavior
			return errors.E500(e, "auth.ServiceImpl.Register.DuplicateUser", "")
		}
		// Failed to register user
		return errors.E500(e, "auth.ServiceImpl.Register.InsertNewUserFail", "")
	}

	return nil
}
