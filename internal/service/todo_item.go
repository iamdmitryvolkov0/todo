package service

import (
	"todo/internal/domain"
	"todo/internal/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem) *TodoItemService {
	return &TodoItemService{repo: repo}
}

func (s *TodoItemService) Create(userID int, listID int, item domain.TodoItem) (int, error) {
	_, err := s.listRepo.ListByID(userID, listID)
	if err != nil {
		return 0, err
	}
	return s.repo.Create(listID, item)
}

func (s *TodoItemService) All(userId, listId int) ([]domain.TodoItem, error) {
	return s.repo.All(userId, listId)
}

func (s *TodoItemService) ItemByID(userId, itemId int) (domain.TodoItem, error) {
	return s.repo.ItemByID(userId, itemId)
}

func (s *TodoItemService) Update(userId, itemId int, input domain.UpdateItemInput) error {
	return s.repo.Update(userId, itemId, input)
}

func (s *TodoItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}
