package repository

import (
	"fmt"

	"github.com/SimilarEgs/CRUD-TODO-LIST/internal/entity"
	"github.com/jmoiron/sqlx"
)

type TodoItemRepository struct {
	db *sqlx.DB
}

func NewTodoItemRepository(db *sqlx.DB) *TodoItemRepository {
	return &TodoItemRepository{db: db}
}

func (r *TodoItemRepository) CreateItem(listId int64, input entity.TodoItem) (int64, error) {

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	createItem := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoItemsTable)

	var itemId int64

	row := tx.QueryRow(createItem, input.Title, input.Description)
	err = row.Scan(&itemId)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListsItems := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES ($1, $2)", listsItemsTable)

	_, err = tx.Exec(createListsItems, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (r *TodoItemRepository) GetAllItems(userId, listId int64) ([]entity.TodoItem, error) {

	var todoListItems []entity.TodoItem

	getAllItemsQuery := fmt.Sprintf(`
	SELECT ti.id, ti.title, ti.description, ti.done 
	FROM %s ti INNER JOIN %s li ON li.item_id = ti.id INNER JOIN %s ul ON ul.list_id = li.list_id 
	WHERE li.list_id = $1 AND ul.user_id = $2`,
		todoItemsTable, listsItemsTable, usersListsTable)

	if err := r.db.Select(&todoListItems, getAllItemsQuery, listId, userId); err != nil {
		return nil, err
	}

	return todoListItems, nil
}

func (r *TodoItemRepository) GetItemById(userId, itemId int64) (entity.TodoItem, error) {

	var todoListItem entity.TodoItem

	getItemById := fmt.Sprintf(`
	SELECT ti.id, ti.title, ti.description, ti.done
	FROM %s ti INNER JOIN %s li ON li.item_id = ti.id INNER JOIN %s ul ON ul.list_id = li.list_id
	WHERE ti.id = $1 AND ul.user_id = $2`,
		todoItemsTable, listsItemsTable, usersListsTable)

	if err := r.db.Get(&todoListItem, getItemById, itemId, userId); err != nil {
		return todoListItem, err
	}

	return todoListItem, nil
}
