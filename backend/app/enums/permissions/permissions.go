package permissions

type Permission string

const (
	UsersShow     Permission = "UsersShow"
	UsersCreate   Permission = "UsersCreate"
	UsersUpdate   Permission = "UsersUpdate"
	UsersDelete   Permission = "UsersDelete"
	UsersSetRoles Permission = "UsersSetRoles"
	RolesShow     Permission = "RolesShow"
	RolesCreate   Permission = "RolesCreate"
	RolesUpdate   Permission = "RolesUpdate"
	RolesDelete   Permission = "RolesDelete"
	FilesShow     Permission = "FilesShow"
	FilesCreate   Permission = "FilesCreate"
	FilesUpdate   Permission = "FilesUpdate"
	FilesDelete   Permission = "FilesDelete"
)
