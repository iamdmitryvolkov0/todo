package service

import (
	"todo/internal/domain"
	"todo/internal/repository"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userID int, list domain.TodoList) (int, error)
	All(userID int) ([]domain.TodoList, error)
	ListByID(userID int, listID int) (domain.TodoList, error)
	Update(userID int, listID int, input domain.UpdateListInput) error
	Delete(userID int, listID int) error
}

type TodoItem interface {
	Create(userId, listId int, item domain.TodoItem) (int, error)
	All(userId, listId int) ([]domain.TodoItem, error)
	ItemByID(userId, itemId int) (domain.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input domain.UpdateItemInput) error
}
type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		TodoList:      NewTodoListService(repo.TodoList),
		TodoItem:      NewTodoItemService(repo.TodoItem),
	}
}
