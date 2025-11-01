package database

import (
	"database/sql"
	"galaveg/database/migrations"
)

type Migration struct {
	Uuid string
	Up   func(db *sql.DB) error
	Down func(db *sql.DB) error
}

var Migrations = []Migration{
	Migration{
		Uuid: "0001_01_01_000000_create_users_table",
		Up:   migrations.CreateUsersTable00010101000000Up,
		Down: migrations.CreateUsersTable00010101000000Down,
	},
	Migration{
		Uuid: "2025_10_31_115348_create_roles_table",
		Up:   migrations.CreateRolesTable20251031115348Up,
		Down: migrations.CreateRolesTable20251031115348Down,
	},
}
