package roles

import "time"

type RoleDto struct {
	ID          ID
	Code        string
	Name        string
	Description string
	Permissions []string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
