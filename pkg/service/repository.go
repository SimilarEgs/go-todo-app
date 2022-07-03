package service

import (
	"github.com/SimilarEgs/CRUD-TODO-LIST/internal/entity"
	"github.com/SimilarEgs/CRUD-TODO-LIST/pkg/repository"
	"github.com/jmoiron/sqlx"
)

type RepositoryAuthorization interface {
	CreateUser(user entity.User) (int, error)
}

type RepositoryTodoList interface {
}

type RepositoryTodoItem interface {
}

type Repository struct {
	RepositoryAuthorization
	RepositoryTodoList
	RepositoryTodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		// initializing repository  
		RepositoryAuthorization: repository.NewAuthRepository(db),
	}
}
