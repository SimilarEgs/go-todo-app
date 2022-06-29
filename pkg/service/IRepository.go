package service

import "github.com/jmoiron/sqlx"

type DBAuthorization interface {
}

type DBTodoList interface {
}

type DBTodoItem interface {
}

type Repository struct {
	Authorization
	TodoItem
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
