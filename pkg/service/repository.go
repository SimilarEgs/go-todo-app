package service

import (
	"github.com/SimilarEgs/go-todo-app/internal/entity"
	"github.com/SimilarEgs/go-todo-app/pkg/repository"
	"github.com/jmoiron/sqlx"
)

type RepositoryAuthorization interface {
	CreateUser(user entity.User) (int64, error)
	GetUser(username string) (entity.User, error)
}

type RepositoryTodoList interface {
	CreateList(userId int64, todoList entity.CreateListInput) (int64, error)
	GetAllLists(userId int64) ([]entity.Todolist, error)
	GetListById(userId, listId int64) (entity.Todolist, error)
	DeleteListById(userId, listId int64) error
	UpdateListById(userId, listId int64, input entity.UpdateListInput) error
}

type RepositoryTodoItem interface {
	CreateItem(listId int64, input entity.TodoItem) (int64, error)
	GetAllItems(userId, listId int64) ([]entity.TodoItem, error)
	GetItemById(userId, itemId int64) (entity.TodoItem, error)
	DeleteItemById(userId, itemId int64) error
	UpdateItemById(userId, itemId int64, input entity.UpdateItemInput) error
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
		RepositoryTodoItem:      repository.NewTodoItemRepository(db),
	}
}
