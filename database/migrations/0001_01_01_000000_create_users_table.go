package migrations

import "galaveg/bootstrap/providers"

func CreateUsersTable00010101000000Up(ctx *providers.Context) error {
	query := `CREATE TABLE ` + ctx.Cfg.Db.Prefix + `users (
id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
email VARCHAR(255) NOT NULL UNIQUE,
password VARCHAR(255) CHARACTER SET ascii COLLATE ascii_bin NULL DEFAULT NULL,
locale VARCHAR(6) CHARACTER SET ascii COLLATE ascii_bin NULL DEFAULT NULL,
surname VARCHAR(255) NULL DEFAULT NULL,
name VARCHAR(255) NULL DEFAULT NULL,
patronymic VARCHAR(255) NULL DEFAULT NULL,
is_super_admin BOOLEAN NOT NULL DEFAULT FALSE,
roles_ids JSON NULL DEFAULT NULL,
avatar_id BIGINT UNSIGNED NULL DEFAULT NULL
);`
	_, err := ctx.Db.Exec(query)
	if err != nil {
		return err
	}

	query = `ALTER TABLE ` + ctx.Cfg.Db.Prefix + `users ADD INDEX avatar_idx (avatar_id);`
	_, err = ctx.Db.Exec(query)
	if err != nil {
		return err
	}

	query = `INSERT INTO ` + ctx.Cfg.Db.Prefix + `users (id, email, is_super_admin, roles_ids) VALUES (1, 'admin@admin.example', true, '[1]');`
	_, err = ctx.Db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func CreateUsersTable00010101000000Down(ctx *providers.Context) error {
	query := `DROP TABLE ` + ctx.Cfg.Db.Prefix + `users;`
	_, err := ctx.Db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
