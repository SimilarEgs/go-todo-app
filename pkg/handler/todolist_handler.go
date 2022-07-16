package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/SimilarEgs/go-todo-app/internal/entity"
	"github.com/SimilarEgs/go-todo-app/utils"
	"github.com/gin-gonic/gin"
)

func (h *Hanlder) createList(c *gin.Context) {

	// fetching user ID
	userId, ok := c.Get(userCTX)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "[Error] user id not found")
		return
	}

	// var for storing input data
	var input entity.Todolist

	// validate input data
	if err := c.ShouldBindJSON(&input); err != nil {
		msg := fmt.Sprintf("[Error] invalid request, try again: %v", err)
		newErrorResponse(c, http.StatusBadRequest, msg)
		return
	}

	// calling service layer method
	listId, err := h.services.TodoList.CreateList(userId.(int64), input)
	if err != nil {
		msg := fmt.Sprintf("[Error] operation failed, try again: %v", err)
		newErrorResponse(c, http.StatusInternalServerError, msg)
		return
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
	userLists, err := h.services.TodoList.GetAllLists(userId.(int64))
	if err != nil {
		if err == sql.ErrNoRows {
			newErrorResponse(c, http.StatusInternalServerError, "[Error] there are no todo lists for the current user")
			return
		}
		msg := fmt.Sprintf("[Error] connection error, try again: %v", err)
		newErrorResponse(c, http.StatusInternalServerError, msg)
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

	// fetching listId and convert it into int64
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "[Error] incorect list id")
		return
	}

	// calling service layer method
	userList, err := h.services.TodoList.GetListById(userId.(int64), int64(listId))

	// error handling
	if err != nil {
		if err == sql.ErrNoRows {
			msg := fmt.Sprintf("[Error] todo list with ID %d - not found", listId)
			newErrorResponse(c, http.StatusBadRequest, msg)
			return
		}
		msg := fmt.Sprintf("[Error] connection error, try again: %v", err)
		newErrorResponse(c, http.StatusInternalServerError, msg)
		return
	}

	// if operation was successfully done, send code 200 to the client and json with user todoList
	c.JSON(http.StatusOK, map[string]interface{}{
		"Todolist": userList,
	})

}

func (h *Hanlder) deleteListById(c *gin.Context) {

	// fetching user ID
	userId, ok := c.Get(userCTX)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "[Error] user id not found")
		return
	}

	// fetching listId and convert it into int64
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "[Error] incorect list id")
		return
	}

	// calling service layer method
	err = h.services.TodoList.DeleteListById(userId.(int64), int64(listId))

	// error handling
	if err != nil {
		if err == sql.ErrNoRows {
			msg := fmt.Sprintf("[Error] todo list with ID %d - not found", listId)
			newErrorResponse(c, http.StatusBadRequest, msg)
			return
		}
		// probably code error, cuz error above will allways trown instead of this one
		if err == utils.ErrRowCntDel {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		msg := fmt.Sprintf("[Error] connection error, try again: %v", err)
		newErrorResponse(c, http.StatusInternalServerError, msg)
		return
	}

	// formating response msg
	msg := fmt.Sprintf("[Info] TodoList with ID %d - was successfully deleted", listId)

	// if operation was successfully done, send code 204 to the client and response msg about successful deletion
	c.JSON(http.StatusOK, msg)
}

func (h *Hanlder) updateListById(c *gin.Context) {

	// fetching user ID
	userId, ok := c.Get(userCTX)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "[Error] user id not found")
		return
	}

	// fetching listId and convert it into int64
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "[Error] incorect list id")
		return
	}

	// var for storing input data
	var input entity.UpdateListInput

	// validate input data
	if err := c.BindJSON(&input); err != nil {
		msg := fmt.Sprintf("[Error] invalid request, try again: %v", err)
		newErrorResponse(c, http.StatusBadRequest, msg)
		return
	}

	// calling service layer method
	err = h.services.TodoList.UpdateListById(userId.(int64), int64(listId), input)

	// error handling
	if err != nil {
		if err == utils.ErrRowCntUp {
			msg := fmt.Sprintf("[Error] todo list with ID %d - not found", listId)
			newErrorResponse(c, http.StatusBadRequest, msg)
			return
		}
		if err == utils.ErrEmptyPayload {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		msg := fmt.Sprintf("[Error] connection error, try again: %v", err)
		newErrorResponse(c, http.StatusInternalServerError, msg)
		return
	}

	// formating response msg
	msg := fmt.Sprintf("[Info] TodoList with ID %d - was successfully updated", listId)

	// if operation was successfully done, send code 200 to the client and json with response msg
	c.JSON(http.StatusOK, msg)
}
