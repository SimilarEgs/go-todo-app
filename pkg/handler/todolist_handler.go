package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/SimilarEgs/go-todo-app/internal/entity"
	"github.com/SimilarEgs/go-todo-app/utils"
	"github.com/gin-gonic/gin"
)

// @Summary     CreateList
// @Security    ApiKeyAuth
// @Tags        Lists
// @Description API endpoint of creating a TodoList
// @ID          create-list
// @Accept      json
// @Produce     json
// @Param       input       body      entity.CreateListInput true "TodoList data"
// @Success     201         {integer} integer                1
// @Failure     400         {object}  errorResponse
// @Failure     404         {object}  errorResponse
// @Failure     500         {object}  errorResponse
// @Failure     default     {object}  errorResponse
// @Router      /api/lists/ [post]
func (h *Hanlder) createList(c *gin.Context) {

	// fetching user ID
	userId, ok := c.Get(userCTX)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "[Error] user id not found")
		return
	}

	// var for storing input data
	var input entity.CreateListInput

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

type getAllListsResponse struct {
	Data []entity.Todolist `json:"data"`
}

// @Summary     GetLists
// @Security    ApiKeyAuth
// @Tags        Lists
// @Description API endpoint of getting all TodoLists
// @ID          get-all-Lists
// @Accept      json
// @Produce     json
// @Success     200          {object} getAllListsResponse
// @Failure     400          {object} errorResponse
// @Failure     404          {object} errorResponse
// @Failure     500          {object} errorResponse
// @Failure     default      {object} errorResponse
// @Router      /api/lists/  [get]
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
	c.JSON(http.StatusOK, getAllListsResponse{
		Data: userLists,
	})
}

// @Summary     GetList
// @Security    ApiKeyAuth
// @Tags        Lists
// @Description API endpoint of getting todo list by id
// @ID          get-list-by-id
// @Accept      json
// @Produce     json
// @Param		id			   path		  int true "TodoList ID"
// @Success     200            {object}   entity.Todolist
// @Failure     400            {object}   errorResponse
// @Failure     404            {object}   errorResponse
// @Failure     500            {object}   errorResponse
// @Failure     default        {object}   errorResponse
// @Router      /api/lists/{id} [get]
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
		x := c.Param("id")
		fmt.Println(x)
		log.Println(err)
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

// @Summary     DeleteList
// @Security    ApiKeyAuth
// @Tags        Lists
// @Description API endpoint of deleting TodoList by id
// @ID          delete-list-by-id
// @Accept      json
// @Produce     json
// @Param		id			     path		int	 true "TodoList ID"
// @Success     200              {string}   json
// @Failure     400     		 {object}   errorResponse
// @Failure     404              {object}   errorResponse
// @Failure     500              {object}   errorResponse
// @Failure     default          {object}   errorResponse
// @Router      /api/lists/{id}  [delete]
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

// @Summary     UpdateList
// @Security    ApiKeyAuth
// @Tags        Lists
// @Description API endpoint of updating TodoList by id
// @ID          delete-list-by-id
// @Accept      json
// @Produce     json
// @Param       input       	 body       entity.UpdateListInput true "TodoList update data"
// @Param		id			     path		int	                   true "TodoList ID"
// @Success     200              {string}   json
// @Failure     400              {object}   errorResponse
// @Failure     404 			 {object}   errorResponse
// @Failure     500              {object}   errorResponse
// @Failure     default          {object}   errorResponse
// @Router      /api/lists/{id}  [put]
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
