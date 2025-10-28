package models

import (
	"galaveg/app/enums/permissions"
	"gorm.io/gorm"
	"time"
)

type Role struct {
	gorm.Model
	ID          uint64 `gorm:"primaryKey"`
	Code        string
	Name        string
	Description *string
	Permissions *[]permissions.Permission
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
