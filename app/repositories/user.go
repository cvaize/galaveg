package repositories

import (
	"database/sql"
	"galaveg/config"
)

type UserRepo struct {
	cfg *config.Config
	db  *sql.DB
}

func (r *UserRepo) CreateByEmail(email string) error {
	query := `INSERT INTO ` + r.cfg.Db.Prefix + `users (email) VALUES (?);`
	_, err := r.db.Exec(query, email)
	if err != nil {
		return err
	}

	return nil
}
