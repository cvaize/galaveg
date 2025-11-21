package repositories

import (
	"database/sql"
	"galaveg/config"
)

type UserRepo struct {
	C  *config.Config
	DB *sql.DB
}

func (r *UserRepo) CreateByEmail(email string) error {
	query := `INSERT INTO ` + r.C.Db.Prefix + `users (email) VALUES (?);`
	_, err := r.DB.Exec(query, email)
	if err != nil {
		return err
	}

	return nil
}
