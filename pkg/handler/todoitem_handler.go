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

// @Summary     CreateItem
// @Security    ApiKeyAuth
// @Tags        TodoListItems
// @Description API endpoint of creating a TodoListItem
// @ID          create-item
// @Accept      json
// @Produce     json
// @Param       input       		   body       entity.CreateItemInput true "Item data"
// @Param       id					   path       integer true "TodoList ID"
// @Success     201         		   {object}   integer                1
// @Failure     400         		   {object}   errorResponse
// @Failure     404         		   {object}   errorResponse
// @Failure     500         		   {object}   errorResponse
// @Failure     default     		   {object}   errorResponse
// @Router      /api/lists/{id}/items/ [post]
func (h *Hanlder) createItem(c *gin.Context) {

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "[Error] invalid request id")
		return
	}

	userId, ok := c.Get(userCTX)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "[Error] user id not found")
		return
	}

	var input entity.TodoItem

	if err := c.BindJSON(&input); err != nil {
		msg := fmt.Sprintf("[Error] invalid request, try again: %v", err)
		newErrorResponse(c, http.StatusBadRequest, msg)
		return
	}

	itemId, err := h.services.TodoItem.CreateItem(userId.(int64), int64(listId), input)
	if err != nil {
		if err == sql.ErrNoRows {
			msg := fmt.Sprintf("[Error] todo list with ID %d - not found", listId)
			newErrorResponse(c, http.StatusBadRequest, msg)
			return
		}

		msg := fmt.Sprintf("[Error] operation failed, try again: %v", err)
		newErrorResponse(c, http.StatusInternalServerError, msg)
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": itemId,
	})

}

type getAllItemsResponse struct {
	Data []entity.TodoItem `json:"data"`
}

// @Summary     GetAllItems
// @Security    ApiKeyAuth
// @Tags        TodoListItems
// @Description API endpoint of getting all TodoListItems
// @ID          get-items
// @Accept      json
// @Produce     json
// @Param       id		       		   path       integer true "TodoList ID"
// @Success     200         		   {object}   getAllItemsResponse
// @Failure     400         		   {object}   errorResponse
// @Failure     404         		   {object}   errorResponse
// @Failure     500         		   {object}   errorResponse
// @Failure     default     		   {object}   errorResponse
// @Router      /api/lists/{id}/items/ [get]
func (h *Hanlder) getAllItems(c *gin.Context) {

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "[Error] invalid request id")
		return
	}

	userId, ok := c.Get(userCTX)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "[Error] user id not found")
		return
	}

	items, err := h.services.TodoItem.GetAllItems(userId.(int64), int64(listId))
	if err != nil {
		if err == sql.ErrNoRows {
			msg := fmt.Sprintf("[Error] todo list with ID %d - does not contain items", listId)
			newErrorResponse(c, http.StatusBadRequest, msg)
			return
		}
		if err == utils.ErrRowCntGet {
			newErrorResponse(c, http.StatusInternalServerError, "[Error] there are no active schedules for current todo list")
			return
		}
		msg := fmt.Sprintf("[Error] connection error, try again: %v", err)
		newErrorResponse(c, http.StatusInternalServerError, msg)
		return
	}

	c.JSON(http.StatusOK, getAllItemsResponse{
		Data: items,
	})
}

// @Summary     GetItemById
// @Security    ApiKeyAuth
// @Tags        TodoListItems
// @Description API endpoint of getting TodoListItem by ID
// @ID          get-item
// @Accept      json
// @Produce     json
// @Param       id		       		   path       integer true "TodoItem ID"
// @Success     200         		   {object}   entity.TodoItem
// @Failure     400         		   {object}   errorResponse
// @Failure     404         		   {object}   errorResponse
// @Failure     500         		   {object}   errorResponse
// @Failure     default     		   {object}   errorResponse
// @Router      /api/items/{id}        [get]
func (h *Hanlder) getItemById(c *gin.Context) {

	itemId, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "[Error] invalid request id")
		return
	}

	userId, ok := c.Get(userCTX)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "[Error] user id not found")
		return
	}

	item, err := h.services.TodoItem.GetItemById(userId.(int64), int64(itemId))
	if err != nil {
		if err == sql.ErrNoRows {
			msg := fmt.Sprintf("[Error] todo list item with ID %d - not found", itemId)
			newErrorResponse(c, http.StatusBadRequest, msg)
			return
		}
		msg := fmt.Sprintf("[Error] connection error, try again: %v", err)
		newErrorResponse(c, http.StatusInternalServerError, msg)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"TodoList Task": item,
	})
}

// @Summary     DeleteItemById
// @Security    ApiKeyAuth
// @Tags        TodoListItems
// @Description API endpoint of deleting TodoListItem by ID
// @ID          delete-item
// @Accept      json
// @Produce     json
// @Param       id		       		   path       integer true "TodoItem ID"
// @Success     200              	   {string}   json
// @Failure     400         		   {object}   errorResponse
// @Failure     404         		   {object}   errorResponse
// @Failure     500         		   {object}   errorResponse
// @Failure     default     		   {object}   errorResponse
// @Router      /api/items/{id}        [delete]
func (h *Hanlder) deleteItemById(c *gin.Context) {

	itemId, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "[Error] invalid request id")
		return
	}

	userId, ok := c.Get(userCTX)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "[Error] user id not found")
		return
	}

	err = h.services.TodoItem.DeleteItemById(userId.(int64), int64(itemId))

	if err != nil {
		if err == sql.ErrNoRows {
			msg := fmt.Sprintf("[Error] todo list item with ID - %d not found", itemId)
			newErrorResponse(c, http.StatusBadRequest, msg)
			return
		}
		if err == utils.ErrRowCntDel {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		msg := fmt.Sprintf("[Error] connection error, try again: %v", err)
		newErrorResponse(c, http.StatusInternalServerError, msg)
		return
	}

	msg := fmt.Sprintf("[Info] todo list item with ID %d - was successfully deleted", itemId)
	c.JSON(http.StatusOK, msg)

}

// @Summary     UpdateItemById
// @Security    ApiKeyAuth
// @Tags        TodoListItems
// @Description API endpoint of updating TodoListItem by ID
// @ID          update-item
// @Accept      json
// @Produce     json
// @Param       input		       	   body       entity.UpdateItemInput true "TodoItem ID"
// @Param       id		       		   path       integer 				 true "TodoItem ID"
// @Success     200              	   {string}   json
// @Failure     400         		   {object}   errorResponse
// @Failure     404         		   {object}   errorResponse
// @Failure     500         		   {object}   errorResponse
// @Failure     default     		   {object}   errorResponse
// @Router      /api/items/{id}        [put]
func (h *Hanlder) updateItemById(c *gin.Context) {

	itemId, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "[Error] invalid request id")
		return
	}

	userId, ok := c.Get(userCTX)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "[Error] user id not found")
		return
	}

	var input entity.UpdateItemInput

	if err := c.BindJSON(&input); err != nil {
		msg := fmt.Sprintf("[Error] invalid request, try again: %v", err)
		newErrorResponse(c, http.StatusBadRequest, msg)
		return
	}

	err = h.services.TodoItem.UpdateItemById(userId.(int64), int64(itemId), input)

	if err != nil {
		if err == utils.ErrRowCntUp {
			msg := fmt.Sprintf("[Error] todo list item with ID  - %d not found", itemId)
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

	msg := fmt.Sprintf("[Info] todo list item with ID - %d was successfully updated", itemId)

	c.JSON(http.StatusOK, msg)
}
