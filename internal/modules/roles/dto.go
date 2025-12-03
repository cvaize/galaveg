package roles

type RoleDto struct {
	ID          ID
	Code        string
	Name        string
	Description string
	Permissions []string
}
