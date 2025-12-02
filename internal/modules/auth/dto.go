package auth

import (
	"galaveg/internal/modules/users"
)

type UserDto struct {
	ID           users.UserID
	Email        EmailVO
	PasswordHash string
}
