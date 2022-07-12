package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/SimilarEgs/go-todo-app/internal/entity"
	"github.com/SimilarEgs/go-todo-app/utils"
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
		return nil, err // sql.ErrNoRows
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

	if len(todoListItems) == 0 {
		return nil, utils.ErrRowCntGet
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

	var mock entity.TodoItem

	getItemById := fmt.Sprintf(`
	SELECT ti.id, ti.title, ti.description, ti.done
	FROM %s ti INNER JOIN %s li ON li.item_id = ti.id INNER JOIN %s ul ON ul.list_id = li.list_id
	WHERE ti.id = $1 AND ul.user_id = $2`,
		todoItemsTable, listsItemsTable, usersListsTable)

	err := r.db.Get(&mock, getItemById, itemId, userId)
	if err != nil {
		return err
	}

	deleteItemById := fmt.Sprintf(`
	DELETE FROM %s ti USING %s li, %s ul
	WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $1 AND ti.id = $2`,
		todoItemsTable, listsItemsTable, usersListsTable)

	res, err := r.db.Exec(deleteItemById, userId, itemId)

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return utils.ErrRowCntDel
	}

	return err
}

func (r *TodoItemRepository) UpdateItemById(userId, itemId int64, input entity.UpdateItemInput) error {

	var mock entity.TodoItem

	getItemById := fmt.Sprintf(`
	SELECT ti.id, ti.title, ti.description, ti.done
	FROM %s ti INNER JOIN %s li ON li.item_id = ti.id INNER JOIN %s ul ON ul.list_id = li.list_id
	WHERE ti.id = $1 AND ul.user_id = $2`,
		todoItemsTable, listsItemsTable, usersListsTable)

	err := r.db.Get(&mock, getItemById, itemId, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return utils.ErrRowCntUp
		}
		return err
	}

	holdValues := make([]string, 0)
	args := make([]interface{}, 0)
	argsId := 1

	if input.Title != nil {
		holdValues = append(holdValues, fmt.Sprintf("title=$%d", argsId))
		args = append(args, *input.Title)
		argsId++
	}
	if input.Description != nil {
		holdValues = append(holdValues, fmt.Sprintf("description=$%d", argsId))
		args = append(args, *input.Description)
		argsId++
	}
	if input.Done != nil {
		holdValues = append(holdValues, fmt.Sprintf("done=$%d", argsId))
		args = append(args, *input.Done)
		argsId++
	}

	setQuery := strings.Join(holdValues, ",")

	updateItemById := fmt.Sprintf(`
	UPDATE %s ti SET %s FROM %s li, %s ul
	WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id =$%d AND ti.id =$%d`,
		todoItemsTable, setQuery, listsItemsTable, usersListsTable, argsId, argsId+1)

	args = append(args, userId, itemId)

	_, err = r.db.Exec(updateItemById, args...)

	return err
}
