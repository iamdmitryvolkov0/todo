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

type TodoList interface {
	Create(userID int, list domain.TodoList) (int, error)
	All(userID int) ([]domain.TodoList, error)
	ListByID(userID int, listID int) (domain.TodoList, error)
	Update(userID int, listID int, input domain.UpdateListInput) error
	Delete(userID int, listID int) error
}

type TodoItem interface {
	Create(listID int, item domain.TodoItem) (int, error)
	All(userID int, listID int) ([]domain.TodoItem, error)
	ItemByID(userID int, itemID int) (domain.TodoItem, error)
	Update(userID int, itemID int, input domain.UpdateItemInput) error
	Delete(userID int, itemID int) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		TodoList:      NewTodoListRepository(db),
		TodoItem:      NewToDoItemRepository(db),
	}
}
