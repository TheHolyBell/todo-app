package service

import (
	"rest_api"
	"rest_api/pkg/repository"
)

type Authorization interface {
	CreateUser(user rest_api.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list rest_api.TodoList) (int, error)
	GetAll(userId int) ([]rest_api.TodoList, error)
	GetById(userId, id int) (rest_api.TodoList, error)
	Update(userId, id int, input rest_api.UpdateListInput) error
	Delete(userId, id int) error
}

type TodoItem interface {
	Create(userId, listId int, item rest_api.TodoItem) (int, error)
	GetAll(userId, listId int) ([]rest_api.TodoItem, error)
	GetById(userId, itemId int) (rest_api.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input rest_api.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem),
	}
}
