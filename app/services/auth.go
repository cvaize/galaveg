package services

import (
	"galaveg/app/dto"
	"galaveg/utils/logger"
	"strings"
)

type AuthService struct {
	US *UserService
	TS *TranslatorService
	HS *HashService
}

func (s *AuthService) Login(email, password string) (dto.UserID, error) {

	return 0, nil
}

func (s *AuthService) Register(email, password string) (dto.UserID, *dto.Error) {
	if exists, e := s.US.ExistsByEmail(email); exists {
		logger.Infof("AuthService.Register <- UserService.ExistsByEmail: %v", e)
		return 0, dto.NewBadRequest("AuthService.DuplicateEmail", e.Error())
	}

	password, e := s.EncodePassword(password)
	if e != nil {
		logger.Fatalf("AuthService.Register <- AuthService.EncodePassword: %v", e)
		return 0, dto.NewInternal("AuthService.Fail", e.Error())
	}

	userId, err := s.US.Create(&dto.User{Email: email, Password: password})
	if err != nil {
		eStr := err.Error()
		if strings.Contains(eStr, "Duplicate entry") {
			if strings.Contains(eStr, ".email'") {
				logger.Infof("AuthService.Register <- UserService.CreateByEmail: %v", err)
				return 0, dto.NewBadRequest("AuthService.DuplicateEmail", eStr)
			}
			logger.Fatalf("AuthService.Register <- UserService.CreateByEmail: %v", err)
			return 0, dto.NewBadRequest("AuthService.Duplicate", eStr)
		}
		logger.Fatalf("AuthService.Register <- UserService.CreateByEmail: %v", err)
		return 0, dto.NewBadRequest("AuthService.InsertNewUserFail", eStr)
	}

	return userId, nil
}

func (s *AuthService) EncodePassword(password string) (string, error) {
	password, e := s.HS.HashPassword(password)
	if e != nil {
		logger.Fatalf("AuthService.EncodePassword <- HashService.HashPassword: %v", e)
		return password, dto.NewInternal("AuthService.Fail", e.Error())
	}

	return password, nil
}

func (s *AuthService) VerifyPassword(password, hash string) (bool, error) {
	is, e := s.HS.VerifyPassword(password, hash)
	if e != nil {
		logger.Fatalf("AuthService.VerifyPassword <- HashService.VerifyPassword: %v", e)
		return is, dto.NewInternal("AuthService.Fail", e.Error())
	}

	return is, nil
}
