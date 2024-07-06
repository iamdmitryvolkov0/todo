package repository

import (
	"github.com/jmoiron/sqlx"
	"todo/internal/domain"
)

const (
	UsersTable      = "users"
	TodoListsTable  = "todo_lists"
	UsersListsTable = "users_lists"
	TodoItemsTable  = "todo_items"
	ListsItemsTable = "lists_items"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GetUser(username, password string) (domain.User, error)
}

type ToDoList interface {
}

type ToDoItem interface {
}

type Repository struct {
	Authorization
	ToDoList
	ToDoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
	}
}
