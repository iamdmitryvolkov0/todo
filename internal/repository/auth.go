package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"todo/internal/domain"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user domain.User) (int, error) {
	var id int
	query := fmt.Sprintf(
		"INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id",
		UsersTable,
	)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthRepository) GetUser(username, password string) (domain.User, error) {
	var user domain.User
	query := fmt.Sprintf(
		"SELECT id FROM %s WHERE username = $1 AND password_hash = $2",
		UsersTable,
	)

	err := r.db.Get(&user, query, username, password)

	return user, err
}
