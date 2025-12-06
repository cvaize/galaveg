package auth

import (
	"database/sql"
	"errors"
	"galaveg/internal/infrastructures/db"
	moduleErrors "galaveg/internal/modules/errors"
)

type DbRepo = *DbRepoImpl

type DbRepoImpl struct {
	db         db.Db
	usersTable string
}

type DbRepoImplSettings struct {
	Db         db.Db
	UsersTable string
	Prefix     string
}

func NewDbRepoImpl(settings DbRepoImplSettings) *DbRepoImpl {
	if settings.UsersTable == "" {
		settings.UsersTable = "users"
	}
	usersTable := settings.Prefix + settings.UsersTable
	return &DbRepoImpl{settings.Db, usersTable}
}

// FirstByEmail retrieves a user by email from the database
func (r *DbRepoImpl) FirstByEmail(email string) (*UserDto, *moduleErrors.Error) {
	var user UserDto
	query := "SELECT id, email, password FROM " + r.usersTable + " WHERE email = ?"
	err := r.db.QueryRow(query, email).Scan(
		&user.Id,
		&user.Email.Value,
		&user.PasswordHash,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Row not found
		}
		return nil, moduleErrors.E500(err, "auth.DbRepoImpl.FirstByEmail.QueryRow", "")
	}
	return &user, nil
}

// Create creates a new user in the database
func (r *DbRepoImpl) Create(user *UserDto) *moduleErrors.Error {
	query := "INSERT INTO " + r.usersTable + " (email, password) VALUES (?, ?)"
	_, err := r.db.Exec(
		query,
		user.Email,
		user.PasswordHash,
	)
	if err != nil {
		return moduleErrors.E500(err, "auth.DbRepoImpl.Create.Exec", "")
	}

	return nil
}
