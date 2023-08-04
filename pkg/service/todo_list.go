package service

import (
	"rest_api"
	"rest_api/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list rest_api.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]rest_api.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId, id int) (rest_api.TodoList, error) {
	return s.repo.GetById(userId, id)
}

func (s *TodoListService) Update(userId, id int, input rest_api.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return nil
	}
	return s.repo.Update(userId, id, input)
}

func (s *TodoListService) Delete(userId, id int) error {
	return s.repo.Delete(userId, id)
}
