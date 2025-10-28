package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID           uint64 `gorm:"primaryKey"`
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
