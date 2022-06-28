package service

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

func NewRepository() *Repository {
	return &Repository{}
}
