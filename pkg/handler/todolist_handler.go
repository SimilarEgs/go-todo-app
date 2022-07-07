package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/SimilarEgs/CRUD-TODO-LIST/internal/entity"
	"github.com/gin-gonic/gin"
)

func (h *Hanlder) createList(c *gin.Context) {

	// fetching user ID
	userId, ok := c.Get(userCTX)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "[Error] user id not found")
		return
	}

	// var for storing user input data
	var input entity.Todolist

	// validate input data
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "[Error] invalid request, try again")
		return
	}

	// calling service layer method
	listId, err := h.services.TodoList.CreateList(userId.(int64), input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "[Error] operation failed, try again")
	}

	// if operation was successfully done, send code 201 to the client and json with id of the created list
	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": listId,
	})

}

func (h *Hanlder) getAllLists(c *gin.Context) {

	// fetching user ID
	userId, ok := c.Get(userCTX)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "[Error] user id not found")
		return
	}

	// calling service layer method
	userLists, err := h.services.GetAllLists(userId.(int64))
	if err != nil {
		if err == sql.ErrNoRows {
			newErrorResponse(c, http.StatusInternalServerError, "[Error] there are no todo lists for the current user")
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, "[Error] connection error, try again")
		return
	}

	// if operation was successfully done, send code 200 to the client and json with slice of user todoLists
	c.JSON(http.StatusOK, map[string]interface{}{
		"TodoLists:": userLists,
	})
}

func (h *Hanlder) getListById(c *gin.Context) {

	// fetching user ID
	userId, ok := c.Get(userCTX)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "[Error] user id not found")
		return
	}

	// fetching listId and it into int64
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "[Error] incorect list id")
		return
	}

	// calling service layer method
	userList, err := h.services.GetListById(userId.(int64), int64(listId))

	// error handling
	if err != nil {
		if err == sql.ErrNoRows {
			newErrorResponse(c, http.StatusNotFound, "[Error] list with such ID not found")
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, "[Error] connection error, try again")
		return
	}

	// if operation was successfully done, send code 200 to the client and json with user todoList
	c.JSON(http.StatusOK, map[string]interface{}{
		"Todolist": userList,
	})

}

func (h *Hanlder) updateListById(c *gin.Context) {
}
func (h *Hanlder) deleteListById(c *gin.Context) {
}
