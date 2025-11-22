package dto

import (
	"time"
)

type UserID uint64

type User struct {
	ID           UserID
	Email        string
	Password     string
	Locale       string
	Surname      string
	Name         string
	Patronymic   string
	IsSuperAdmin bool
	RolesIds     []uint64
	AvatarId     uint64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Roles        []Role
}
