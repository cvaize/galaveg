package roles

import (
	"time"
)

type Role struct {
	ID          uint64
	Code        string
	Name        string
	Description string
	Permissions []string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
