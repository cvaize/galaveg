package services

import "database/sql"

type UserID uint64

type UserService struct {
	DB *sql.DB
}

func (s *UserService) CreateByEmail(email string) {

}
