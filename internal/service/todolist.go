package service

import (
	"todo/internal/domain"
	"todo/internal/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userID int, list domain.TodoList) (int, error) {
	return s.repo.Create(userID, list)
}

func (s *TodoListService) All(userID int) ([]domain.TodoList, error) {
	return s.repo.All(userID)
}

func (s *TodoListService) ListByID(userID int, listID int) (domain.TodoList, error) {
	return s.repo.ListByID(userID, listID)
}

func (s *TodoListService) Delete(userID int, listID int) error {
	return s.repo.Delete(userID, listID)
}
