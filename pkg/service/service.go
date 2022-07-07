package service

import (
	"github.com/SimilarEgs/CRUD-TODO-LIST/internal/entity"
)

type Authorization interface {
	// this method takes User entity as args
	// and return ID of created user in DB
	CreateUser(user entity.User) (int64, error)

	// this  method takes account data as args
	// and return generated JWT
	GenerateToken(username, password string) (string, error)

	// this method takes auth token as args
	// and return ID of the user affter succsessfull parsing
	ParseToken(token string) (int64, error)
}

type TodoList interface {
	// this method takes ID of the user and TodoList entity
	// and return id of created TodoList in db
	CreateList(userId int64, list entity.Todolist) (int64, error)

	// this method takes user ID
	// and return all lists that this user have
	GetAllLists(userId int64) ([]entity.Todolist, error)
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.RepositoryAuthorization),
		TodoList:      NewTodoListService(repos.RepositoryTodoList),
	}
}
