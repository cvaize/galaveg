package migrations

import (
	"galaveg/config"
	"galaveg/connections"
)

func CreateRolesTable00010101000010Up() error {

	query := `CREATE TABLE ` + config.Config.Db.Prefix + `roles (
id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
code VARCHAR(255) NOT NULL UNIQUE,
name VARCHAR(255) NOT NULL UNIQUE,
description VARCHAR(255) NULL DEFAULT NULL,
permissions JSON NULL DEFAULT NULL
);`
	_, err := connections.DB.Exec(query)
	if err != nil {
		return err
	}

	query = `INSERT INTO ` + config.Config.Db.Prefix + `roles (id, code, name, permissions) VALUES (1, 'admin', 'Admin', '["users_show",
"users_create",
"users_update",
"users_delete",
"users_set_roles",
"roles_show",
"roles_create",
"roles_update",
"roles_delete"]');`
	_, err = connections.DB.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func CreateRolesTable00010101000010Down() error {

	query := `DROP TABLE ` + config.Config.Db.Prefix + `roles;`
	_, err := connections.DB.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
