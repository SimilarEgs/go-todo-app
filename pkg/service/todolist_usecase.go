package service

import "github.com/SimilarEgs/CRUD-TODO-LIST/internal/entity"

type TodoListService struct {
	repo RepositoryTodoList
}

func NewTodoListService(repo RepositoryTodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

// this function pass user data at the repository layer, no additional logic is required in the implementation here
func (s *TodoListService) CreateList(userId int64, list entity.Todolist) (int64, error) {
	return s.repo.CreateList(userId, list)
}

// this function pass user ID at the repository layer, no additional logic is required in the implementation here
func (s *TodoListService) GetAllLists(userId int64) ([]entity.Todolist, error) {
	return s.repo.GetAllLists(userId)
}
