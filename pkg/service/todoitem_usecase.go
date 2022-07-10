package service

import "github.com/SimilarEgs/CRUD-TODO-LIST/internal/entity"

type TodoItemService struct {
	repo     RepositoryTodoItem
	listRepo RepositoryTodoList
}

func NewTodoItemService(repo RepositoryTodoItem, listRepo RepositoryTodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) CreateItem(userId, listId int64, input entity.TodoItem) (int64, error) {
	// before send data on repo layer
	// check if TodoList exists and if it belongs to the user
	_, err := s.listRepo.GetListById(userId, listId)
	if err != nil {
		return 0, err
	}

	return s.repo.CreateItem(listId, input)
}

func (s *TodoItemService) GetAllItems(userId, listId int64) ([]entity.TodoItem, error) {
	return s.repo.GetAllItems(userId, listId)
}

func (s *TodoItemService) GetItemById(userId, listId int64) (entity.TodoItem, error) {
	return s.repo.GetItemById(userId, listId)
}

func (s *TodoItemService) DeleteItemById(userId, itemId int64) error {
	return s.repo.DeleteItemById(userId, itemId)
}
