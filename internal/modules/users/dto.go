package users

import (
	"galaveg/internal/modules/roles"
	"time"
)

type UserDto struct {
	ID           ID
	Email        string
	Locale       string
	Surname      string
	Name         string
	Patronymic   string
	IsSuperAdmin bool
	RolesIds     []int64
	AvatarId     int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Roles        []roles.RoleDto
}
