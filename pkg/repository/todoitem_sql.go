package repository

import (
	"database/sql"
	"fmt"

	"github.com/SimilarEgs/CRUD-TODO-LIST/internal/entity"
	"github.com/SimilarEgs/CRUD-TODO-LIST/utils"
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
	
	// before sending query to delete list element in the db
	// checking for list existence -> if not return corresponding error
	var mock entity.Todolist

	getListById := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul ON tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2",
		todoListsTable, usersListsTable)

	err := r.db.Get(&mock, getListById, userId, listId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

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

func (r *TodoItemRepository) DeleteItemById(userId, itemId int64) error {

	deleteItemById := fmt.Sprintf(`
	DELETE FROM %s ti USING %s li, %s ul
	WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $1 AND ti.id = $2`,
		todoItemsTable, listsItemsTable, usersListsTable)

	res, err := r.db.Exec(deleteItemById, userId, itemId)

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil
	}

	if rowsAffected != 1 {
		return utils.ErrRowCnt
	}

	return err
}
