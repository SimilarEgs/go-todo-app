package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/SimilarEgs/go-todo-app/internal/entity"
	"github.com/SimilarEgs/go-todo-app/utils"
	"github.com/jmoiron/sqlx"
)

type TodoListRepository struct {
	db *sqlx.DB
}

func NewTodoListRepository(db *sqlx.DB) *TodoListRepository {
	return &TodoListRepository{db: db}
}

// this methid deals with a transaction of 2 tables:
// 1. insert into todo_lists ...
// 2. insert into users_lists ... (this table links users to their lists)
func (r *TodoListRepository) CreateList(userId int64, todoList entity.Todolist) (int64, error) {

	// prepare new transaction
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	// query for todoLists entry
	createList := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)

	// executing first sql statement
	row := tx.QueryRow(createList, todoList.Title, todoList.Description)

	// var for storing ID of the created list
	var listId int64

	// storing list ID, if any error aborts the transaction
	if err := row.Scan(&listId); err != nil {
		tx.Rollback()
		return 0, err
	}

	// sql query for usersLists entry
	createUserList := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)

	// second sql statement execution
	_, err = tx.Exec(createUserList, userId, listId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return listId, tx.Commit()
}

func (r *TodoListRepository) GetAllLists(userId int64) ([]entity.Todolist, error) {

	// var for storing user todolists
	var todoLists []entity.Todolist

	// sql query for getting all lists with associated user ID
	getAllLists := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul ON tl.id = ul.list_id WHERE ul.user_id = $1",
		todoListsTable, usersListsTable)

	// exec query
	err := r.db.Select(&todoLists, getAllLists, userId)

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

	// probably it's a bad idea, check existence of row, before send query of deletion
	// i guess more idiomatic way, it's to only check affected rows, to avoid DB load

	// mock for checking if a row exists
	var mock entity.Todolist

	// sql query for getting todolist by ID
	getListById := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul ON tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2",
		todoListsTable, usersListsTable)

	// exec checking query
	err := r.db.Get(&mock, getListById, userId, listId)

	// returns an error if there is no list with requested ID
	if err != nil {
		return err
	}

	// sql query for deleting todolist by ID
	deleteListById := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id=$1 AND ul.list_id=$2",
		todoListsTable, usersListsTable)

	// exec query
	res, err := r.db.Exec(deleteListById, userId, listId)

	// check affected rows
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	// if now rows affected return coresponding error
	if rowsAffected != 1 {
		return utils.ErrRowCntDel
	}

	return err
}

// this method checks request data before sending the query to db
// depending on the updated data -> build an SQL statement using preparation args
func (r *TodoListRepository) UpdateListById(userId, listId int64, input entity.UpdateListInput) error {

	// mock to check if the desired update row exists
	var mock entity.Todolist

	// sql query for getting todolist by ID
	getListById := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul ON tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2",
		todoListsTable, usersListsTable)

	// exec checking query
	err := r.db.Get(&mock, getListById, userId, listId)

	// return error if there is no list with requested ID
	if err != nil {
		if err == sql.ErrNoRows {
			return utils.ErrRowCntUp
		}
		return err
	}

	// args preparation
	holdValues := make([]string, 0)
	args := make([]interface{}, 0)
	argsId := 1

	// checking input data
	// if fields are not nill -> append holdValues slice with corresponding placeholder
	// and add input title to the args slice
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

	setQuery := strings.Join(holdValues, ",")

	// sql query to update required fields
	updateListById := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		todoListsTable, setQuery, usersListsTable, argsId, argsId+1)

	args = append(args, listId, userId)

	_, err = r.db.Exec(updateListById, args...)

	return err
}
