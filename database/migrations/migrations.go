package migrations

import "galaveg/internal/bootstrap/http/context"

type Migration struct {
	Uuid string
	Up   func(ctx *context.Context) error
	Down func(ctx *context.Context) error
}

func GetMigrations() []Migration {
	migrations := []Migration{
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
		// NEW_MIGRATIONS_TAG: MustConfig migrations are added here automatically by the "make:migration" command.
	}

	return migrations
}
