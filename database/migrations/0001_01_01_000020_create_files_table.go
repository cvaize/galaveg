package migrations

import "galaveg/internal/bootstrap/http/context"

func CreateFilesTable00010101000020Up(ctx *context.Context) error {
	query := `CREATE TABLE ` + ctx.Cfg.Db.Prefix + `files (
   id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
   filename VARCHAR(2048) CHARACTER SET ascii COLLATE ascii_bin NOT NULL COMMENT 'The file name is made up of the hash, size, and extensions obtained when uploading the file, by mask: [hash]-[size].[extensions].',
   path VARCHAR(2048) CHARACTER SET ascii COLLATE ascii_bin NOT NULL COMMENT 'The path where the file is saved on disk.',
   mime VARCHAR(255) CHARACTER SET ascii COLLATE ascii_bin NULL DEFAULT NULL COMMENT 'The file type.',
   hash VARCHAR(64) CHARACTER SET ascii COLLATE ascii_bin NULL DEFAULT NULL COMMENT 'Hash of the sha256 file.',
   size BIGINT UNSIGNED NULL DEFAULT NULL COMMENT 'The file size in bytes.',
   creator_user_id BIGINT UNSIGNED NULL DEFAULT NULL COMMENT 'The first user to upload the file.',
   created_at DATETIME NULL DEFAULT NULL COMMENT 'The datetime of the file creation.',
   updated_at DATETIME NULL DEFAULT NULL COMMENT 'The datetime of the last file update.',
   delete_at DATETIME NULL DEFAULT NULL COMMENT 'After this time, the file must be deleted.',
   deleted_at DATETIME NULL DEFAULT NULL COMMENT 'The datetime when the file was deleted.',
   is_delete BOOLEAN NOT NULL DEFAULT FALSE COMMENT 'Label: whether the file needs to be deleted.',
   is_deleted BOOLEAN NOT NULL DEFAULT FALSE COMMENT 'Label: whether the file has been deleted.',
   disk VARCHAR(255) CHARACTER SET ascii COLLATE ascii_bin NOT NULL COMMENT 'The disk where the file is stored.'
) COMMENT 'The file table.';`
	_, err := ctx.Infra.Db.Exec(query)
	if err != nil {
		return err
	}

	query = `ALTER TABLE ` + ctx.Cfg.Db.Prefix + `files ADD UNIQUE disk_path_udx (disk, path);`
	_, err = ctx.Infra.Db.Exec(query)
	if err != nil {
		return err
	}

	query = `ALTER TABLE ` + ctx.Cfg.Db.Prefix + `files ADD INDEX creator_user_idx (creator_user_id);`
	_, err = ctx.Infra.Db.Exec(query)
	if err != nil {
		return err
	}

	query = `ALTER TABLE ` + ctx.Cfg.Db.Prefix + `files ADD INDEX path_idx (path);`
	_, err = ctx.Infra.Db.Exec(query)
	if err != nil {
		return err
	}

	query = `ALTER TABLE ` + ctx.Cfg.Db.Prefix + `files ADD INDEX filename_idx (filename);`
	_, err = ctx.Infra.Db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func CreateFilesTable00010101000020Down(ctx *context.Context) error {
	query := `DROP TABLE ` + ctx.Cfg.Db.Prefix + `files;`
	_, err := ctx.Infra.Db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
