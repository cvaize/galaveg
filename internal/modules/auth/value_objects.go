package auth

import "strings"

type EmailVO struct {
	Value string
}

func NewEmailVO(value string) EmailVO {
	value = strings.TrimSpace(strings.ToLower(value))
	return EmailVO{value}
}

type PasswordVO struct {
	Value string
}

func NewPasswordVO(value string) PasswordVO {
	value = strings.TrimSpace(value)
	return PasswordVO{value}
}
