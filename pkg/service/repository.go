package service

import (
	"github.com/SimilarEgs/CRUD-TODO-LIST/internal/entity"
	"github.com/SimilarEgs/CRUD-TODO-LIST/pkg/repository"
	"github.com/jmoiron/sqlx"
)

type RepositoryAuthorization interface {
	CreateUser(user entity.User) (int64, error)
	GetUser(username string) (entity.User, error)
}

type RepositoryTodoList interface {
	CreateList(userId int64, todoList entity.Todolist) (int64, error)
	GetAllLists(userId int64) ([]entity.Todolist, error)
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
		RepositoryTodoList:      repository.NewTodoListRepository(db),
	}
}
