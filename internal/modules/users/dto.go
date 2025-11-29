package users

import (
	"galaveg/internal/modules/roles"
	"time"
)

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
	Roles        []roles.Role
}
