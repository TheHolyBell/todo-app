package repository

import (
	"github.com/jmoiron/sqlx"
	"rest_api"
)

type Authorization interface {
	CreateUser(user rest_api.User) (int, error)
	GetUser(username, password string) (rest_api.User, error)
}

type TodoList interface {
	Create(userId int, list rest_api.TodoList) (int, error)
	GetAll(userId int) ([]rest_api.TodoList, error)
	GetById(userId, id int) (rest_api.TodoList, error)
	Update(userId, id int, input rest_api.UpdateListInput) error
	Delete(userId, id int) error
}

type TodoItem interface {
	Create(listId int, item rest_api.TodoItem) (int, error)
	GetAll(userId, listId int) ([]rest_api.TodoItem, error)
	GetById(userId, itemId int) (rest_api.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input rest_api.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostrgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
