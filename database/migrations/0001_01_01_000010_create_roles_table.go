package migrations

import "galaveg/bootstrap/providers"

func CreateRolesTable00010101000010Up(ctx *providers.Context) error {
	query := `CREATE TABLE ` + ctx.C.Db.Prefix + `roles (
id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
code VARCHAR(255) NOT NULL UNIQUE,
name VARCHAR(255) NOT NULL UNIQUE,
description VARCHAR(255) NULL DEFAULT NULL,
permissions JSON NULL DEFAULT NULL
);`
	_, err := ctx.DB.Exec(query)
	if err != nil {
		return err
	}

	query = `INSERT INTO ` + ctx.C.Db.Prefix + `roles (id, code, name, permissions) VALUES (1, 'admin', 'Admin', '["users_show",
"users_create",
"users_update",
"users_delete",
"users_set_roles",
"roles_show",
"roles_create",
"roles_update",
"roles_delete"]');`
	_, err = ctx.DB.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func CreateRolesTable00010101000010Down(ctx *providers.Context) error {
	query := `DROP TABLE ` + ctx.C.Db.Prefix + `roles;`
	_, err := ctx.DB.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
