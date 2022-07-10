package service

import (
	"github.com/SimilarEgs/CRUD-TODO-LIST/internal/entity"
)

type Authorization interface {
	// this method takes User entity as args
	// and return ID of created user in DB
	CreateUser(user entity.User) (int64, error)

	// this method takes account data as args
	// and return generated JWT
	GenerateToken(username, password string) (string, error)

	// this method takes auth token as args
	// and return ID of the user affter succsessfull parsing
	ParseToken(token string) (int64, error)
}

type TodoList interface {
	// this method takes ID of the user and TodoList entity as args
	// and return id of created TodoList in db
	CreateList(userId int64, list entity.Todolist) (int64, error)

	// this method takes user ID as args
	// and return all lists that this user have
	GetAllLists(userId int64) ([]entity.Todolist, error)

	// this method takes list and user ID as args
	// and return associated list
	GetListById(userId, listId int64) (entity.Todolist, error)

	// this method takes user and list ID as args
	// and return an error
	DeleteListById(userId, listId int64) error

	// this method takes list and user ID with inut data as args
	// and return an error
	UpdateListById(userId, listId int64, input entity.UpdateListInput) error
}

type TodoItem interface {
	// this method takes user and list ID  with TodoItem entity as args
	// and return id of created TodoItem in db
	CreateItem(userId, listId int64, input entity.TodoItem) (int64, error)

	// this method takes user and list ID
	// and returns all items of that list ID this user have
	GetAllItems(userId, listId int64) ([]entity.TodoItem, error)
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.RepositoryAuthorization),
		TodoList:      NewTodoListService(repo.RepositoryTodoList),
		TodoItem:      NewTodoItemService(repo.RepositoryTodoItem, repo.RepositoryTodoList),
	}
}
