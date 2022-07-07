package repository

import (
	"fmt"

	"github.com/SimilarEgs/CRUD-TODO-LIST/internal/entity"
	"github.com/jmoiron/sqlx"
)

type TodoListRepository struct {
	db *sqlx.DB
}

func NewTodoListRepository(db *sqlx.DB) *TodoListRepository {
	return &TodoListRepository{db: db}
}

// this function deals with a transaction of 2 tables:
// 1. insert into todo_lists ...
// 2. insert into users_lists ... (this table links users to their lists)
func (r *TodoListRepository) CreateList(userId int64, todoList entity.Todolist) (int64, error) {

	// prepare new transaction
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	// var for storing the ID of the created list
	var listId int64

	// query for todoLists entry
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)

	// executing first sql statement
	row := tx.QueryRow(createListQuery, todoList.Title, todoList.Description)

	// storing list ID, if any error aborts the transaction
	if err := row.Scan(&listId); err != nil {
		tx.Rollback()
		return 0, nil
	}

	// sql query for usersLists entry
	createUserListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)

	// second sql statement execution
	_, err = tx.Exec(createUserListQuery, userId, listId)
	if err != nil {
		tx.Rollback()
		return 0, nil
	}

	return listId, tx.Commit()
}

func (r *TodoListRepository) GetAllLists(userId int64) ([]entity.Todolist, error) {

	// var for storing user todolists
	var todoLists []entity.Todolist

	// sql query for getting all lists with associated user ID
	getAllListsQuery := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul ON tl.id = ul.list_id WHERE ul.user_id = $1",
		todoListsTable, usersListsTable)

	// exec query
	err := r.db.Select(&todoLists, getAllListsQuery, userId)

	return todoLists, err
}

func (r *TodoListRepository) GetListById(userId, listId int64) (entity.Todolist, error) {

	// var for storing user todolist
	var todoList entity.Todolist

	// sql query for getting todolist by ID
	getListById := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul ON tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2",
		todoListsTable, usersListsTable)

	// exec query
	err := r.db.Get(&todoList, getListById, userId, listId)

	return todoList, err
}

func (r *TodoListRepository) DeleteListById(userId, listId int64) error {

	// sql query for deleting todolist by ID
	deleteListById := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id=$1 AND ul.list_id=$2",
		todoListsTable, usersListsTable)

	// exec query
	_, err := r.db.Exec(deleteListById, userId, listId)

	return err
}
