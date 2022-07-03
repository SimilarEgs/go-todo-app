package service

import (
	"github.com/SimilarEgs/CRUD-TODO-LIST/internal/entity"
)

type Authorization interface {
	// this method takes User struct as args
	// and return ID of created user in DB
	CreateUser(user entity.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoItem
	TodoList
}

func NewService(repos *Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.RepositoryAuthorization),
	}
}
