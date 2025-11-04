package dto

import (
	"galaveg/app/enums/permissions"
	"time"
)

type Role struct {
	ID          uint64
	Code        string
	Name        string
	Description *string
	Permissions *[]permissions.Permission
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
