package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"todo/internal/domain"
	"todo/internal/storage"
)

type Auth struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *Auth {
	return &Auth{db: db}
}

func (r *Auth) CreateUser(user domain.User) (int, error) {
	var id int
	query := fmt.Sprintf(
		"INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id",
		storage.UsersTable,
	)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
