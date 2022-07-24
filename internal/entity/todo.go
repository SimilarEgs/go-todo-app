package entity

import (
	"github.com/SimilarEgs/go-todo-app/utils"
)

type Todolist struct {
	Id          int    `json:"id"          db:"id"`
	Title       string `json:"title"       db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id"          db:"id"`
	Title       string `json:"title"       db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done"        db:"done"`
}

type ListItem struct {
	Id     int
	ListId int
	ItemId int
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

type CreateListInput struct {
	Title       *string `json:"title" binding:"required"`
	Description *string `json:"description"`
}

// input validator for TodoList struct
func (i UpdateListInput) Validator() error {
	if i.Title == nil && i.Description == nil {
		return utils.ErrEmptyPayload
	}
	return nil
}

type UpdateItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

// input validator for TodoItem struct
func (i UpdateItemInput) Validator() error {
	if i.Title == nil && i.Description == nil && i.Done == nil {
		return utils.ErrEmptyPayload
	}
	return nil
}
