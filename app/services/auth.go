package services

import (
	"galaveg/app/dto"
	"strings"
)

type AuthService struct {
	US *UserService
	TS *TranslatorService
	HS *HashService
	ES *ErrorService
}

func (s *AuthService) Login(email, password string) (dto.UserID, *dto.Error) {
	user, e := s.US.FirstByEmail(email)
	if e != nil {
		// Не удалось получить пользователя
		// Failed to get user
		return 0, s.ES.E500(e, "AuthService.Login.FailedToGetUser", "")
	}

	if user == nil {
		// Пользователь не найден
		// User not found
		return 0, s.ES.E404(e, "AuthService.Login.UserNotFound", "")
	}

	is, e := s.HS.VerifyPassword(password, user.Password)
	if e != nil {
		// Проверить пароль не удалось
		// Failed to verify password
		return 0, s.ES.E500(e, "AuthService.Login.FailedToVerifyPassword", "")
	}
	if !is {
		// Пароль не соответствует сохранённому хешу
		// The password does not match the saved hash
		return 0, s.ES.E401(e, "AuthService.Login.Unauthorized", "")
	}

	return user.ID, nil
}

func (s *AuthService) Register(email, password string) (dto.UserID, *dto.Error) {
	exists, e := s.US.ExistsByEmail(email)
	if e != nil {
		// Не удалось проверить наличие пользователя
		// Failed to check for user existence
		return 0, s.ES.E500(e, "AuthService.Register.FailedToGetUser", "")
	}
	if exists {
		// Такой пользователь уже зарегистрирован
		// This user is already registered
		return 0, s.ES.E400(e, "AuthService.Register.DuplicateUser", "")
	}

	password, e = s.HS.HashPassword(password)
	if e != nil {
		// Не удалось сгенерировать хеш пароля
		// Failed to generate password hash
		return 0, s.ES.E500(e, "AuthService.Register.HashPasswordFail", "")
	}

	userId, err := s.US.Create(&dto.User{Email: email, Password: password})
	if err != nil {
		eStr := err.Error()
		if strings.Contains(eStr, "Duplicate entry") {
			if strings.Contains(eStr, ".email'") {
				// Такой пользователь уже есть
				// There is already such a user
				return 0, s.ES.E400(e, "AuthService.Register.DuplicateUser", "")
			}
			// Такой пользователь уже есть, не предусмотренное поведение
			// There is already such a user, not intended behavior
			return 0, s.ES.E500(e, "AuthService.Register.DuplicateUser", "")
		}
		// Не удалось зарегистрировать пользователя
		// Failed to register user
		return 0, s.ES.E500(e, "AuthService.Register.InsertNewUserFail", "")
	}

	return userId, nil
}
