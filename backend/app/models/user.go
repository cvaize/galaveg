package models

import (
	"time"
)

type User struct {
	ID           uint64
	Email        string
	Locale       *string
	Surname      *string
	Name         *string
	Patronymic   *string
	IsSuperAdmin bool
	RolesIds     []uint64
	AvatarId     *uint64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Roles        []Role
}
