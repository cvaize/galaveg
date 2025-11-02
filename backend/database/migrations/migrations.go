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
		Uuid: "0001_01_01_000010_create_roles_table",
		Up:   CreateRolesTable00010101000010Up,
		Down: CreateRolesTable00010101000010Down,
	},
	Migration{
		Uuid: "0001_01_01_000020_create_files_table",
		Up:   CreateFilesTable00010101000020Up,
		Down: CreateFilesTable00010101000020Down,
	},
	Migration{
		Uuid: "0001_01_01_000020_create_users_files_table",
		Up:   CreateUsersFilesTable00010101000030Up,
		Down: CreateUsersFilesTable00010101000030Down,
	},
}
