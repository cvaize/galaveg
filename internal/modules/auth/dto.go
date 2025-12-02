package auth

import (
	"galaveg/internal/modules/users"
)

type UserDto struct {
	ID           users.ID
	Email        EmailVO
	PasswordHash string
}
