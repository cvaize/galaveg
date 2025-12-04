package auth

type UserDto struct {
	Id           int64
	Email        EmailVO
	PasswordHash string
}
