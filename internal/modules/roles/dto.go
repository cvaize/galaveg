package roles

import "time"

type RoleDto struct {
	Id          int64
	Code        string
	Name        string
	Description string
	Permissions []string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
