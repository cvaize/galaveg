package migrations

type Migration struct {
	Uuid string
	Up   func() error
	Down func() error
}

var Migrations = []Migration{
	Migration{
		Uuid: "0001_01_01_000000_create_users_table",
		Up:   CreateUsersTable00010101000000Up,
		Down: CreateUsersTable00010101000000Down,
	},
	Migration{
		Uuid: "2025_10_31_115348_create_roles_table",
		Up:   CreateRolesTable20251031115348Up,
		Down: CreateRolesTable20251031115348Down,
	},
}
