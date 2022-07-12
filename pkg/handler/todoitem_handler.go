package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/SimilarEgs/CRUD-TODO-LIST/internal/entity"
	"github.com/SimilarEgs/CRUD-TODO-LIST/utils"
	"github.com/gin-gonic/gin"
)

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
			msg := fmt.Sprintf("[Error] todo list item with ID %d - not found", listId)
			newErrorResponse(c, http.StatusBadRequest, msg)
			return
		}
		if err == utils.ErrRowCntGet {
			newErrorResponse(c, http.StatusInternalServerError, "[Error] there are no active schedules for current todo lists")
			return
		}
		msg := fmt.Sprintf("[Error] connection error, try again: %v", err)
		newErrorResponse(c, http.StatusInternalServerError, msg)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"TodoList contents": items,
	})
}

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
