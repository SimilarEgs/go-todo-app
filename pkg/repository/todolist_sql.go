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

	// storing list id, if any error aborts the transaction
	if err := row.Scan(&listId); err != nil{
		tx.Rollback()
		return 0, nil
	}

	// qury for usersLists entry
	createUserListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)

	// second sql statement execution
	_, err = tx.Exec(createUserListQuery, userId, listId)
	if err != nil{
		tx.Rollback()
		return 0, nil
	} 
	
	return listId, tx.Commit()
}
