package service

import "github.com/SimilarEgs/go-todo-app/internal/entity"

type TodoListService struct {
	repo RepositoryTodoList
}

func NewTodoListService(repo RepositoryTodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

// this method passes user data at the repository layer, no additional logic is required in the implementation here
func (s *TodoListService) CreateList(userId int64, list entity.Todolist) (int64, error) {
	return s.repo.CreateList(userId, list)
}

// this method passes user ID at the repository layer, no additional logic is required in the implementation here
func (s *TodoListService) GetAllLists(userId int64) ([]entity.Todolist, error) {
	return s.repo.GetAllLists(userId)
}

// this method passes user and list ID  at the repository layer, no additional logic is required in the implementation here
func (s *TodoListService) GetListById(userId, listId int64) (entity.Todolist, error) {
	return s.repo.GetListById(userId, listId)
}

// this method passes user and list ID at the repository layer, no additional logic is required in the implementation here
func (s *TodoListService) DeleteListById(userId, listId int64) error {
	return s.repo.DeleteListById(userId, listId)
}

// this method validates input data and passes requested args at repo layer
func (s *TodoListService) UpdateListById(userId, listId int64, input entity.UpdateListInput) error {

	if err := input.Validator(); err != nil {
		return err
	}
	
	return s.repo.UpdateListById(userId, listId, input)
}
