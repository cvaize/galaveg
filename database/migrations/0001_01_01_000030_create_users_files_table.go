package migrations

import "galaveg/internal/bootstrap/http/context"

func CreateUsersFilesTable00010101000030Up(ctx *context.Context) error {
	query := `CREATE TABLE ` + ctx.Cfg.Db.Prefix + `users_files (
   id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
   file_id BIGINT UNSIGNED NOT NULL COMMENT 'Relation to the files table.',
   filename VARCHAR(255) CHARACTER SET ascii COLLATE ascii_bin NULL DEFAULT NULL COMMENT 'The file name.',
   path VARCHAR(2048) CHARACTER SET ascii COLLATE ascii_bin NULL DEFAULT NULL COMMENT 'The path or url where you can get the file.',
   upload_filename VARCHAR(255) NULL DEFAULT NULL COMMENT 'The filename received during the upload.',
   mime VARCHAR(255) CHARACTER SET ascii COLLATE ascii_bin NULL DEFAULT NULL COMMENT 'The file type received during the upload.',
   user_id BIGINT UNSIGNED NOT NULL COMMENT 'The user who uploaded the file.',
   created_at DATETIME NULL DEFAULT NULL COMMENT 'The datetime of the file creation.',
   updated_at DATETIME NULL DEFAULT NULL COMMENT 'The datetime of the last file update.',
   deleted_at DATETIME NULL DEFAULT NULL COMMENT 'The datetime when the file was deleted.',
   is_deleted BOOLEAN NOT NULL DEFAULT FALSE COMMENT 'Label: whether the file has been deleted.',
   is_public BOOLEAN NOT NULL DEFAULT FALSE COMMENT 'Label: public file or not.',
   disk VARCHAR(255) CHARACTER SET ascii COLLATE ascii_bin NOT NULL COMMENT 'The disk where the file is stored.'
) COMMENT 'Files belonging to users.';`
	_, err := ctx.Infra.Db.Exec(query)
	if err != nil {
		return err
	}

	query = `ALTER TABLE ` + ctx.Cfg.Db.Prefix + `users_files ADD UNIQUE file_user_udx (user_id, file_id);`
	_, err = ctx.Infra.Db.Exec(query)
	if err != nil {
		return err
	}

	query = `ALTER TABLE ` + ctx.Cfg.Db.Prefix + `users_files ADD INDEX user_idx (user_id);`
	_, err = ctx.Infra.Db.Exec(query)
	if err != nil {
		return err
	}

	query = `ALTER TABLE ` + ctx.Cfg.Db.Prefix + `users_files ADD INDEX path_idx (path);`
	_, err = ctx.Infra.Db.Exec(query)
	if err != nil {
		return err
	}

	query = `ALTER TABLE ` + ctx.Cfg.Db.Prefix + `users_files ADD INDEX filename_idx (filename);`
	_, err = ctx.Infra.Db.Exec(query)
	if err != nil {
		return err
	}

	query = `ALTER TABLE ` + ctx.Cfg.Db.Prefix + `users_files ADD INDEX upload_filename_idx (upload_filename);`
	_, err = ctx.Infra.Db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func CreateUsersFilesTable00010101000030Down(ctx *context.Context) error {
	query := `DROP TABLE ` + ctx.Cfg.Db.Prefix + `users_files;`
	_, err := ctx.Infra.Db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
