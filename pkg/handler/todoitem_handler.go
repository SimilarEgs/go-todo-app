package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/SimilarEgs/CRUD-TODO-LIST/internal/entity"
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
			newErrorResponse(c, http.StatusNotFound, "[Error] todo list with such ID not found")
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
		log.Println(err)
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
			newErrorResponse(c, http.StatusInternalServerError, "[Error] todo task does not exist yet, return no rows")
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

func (h *Hanlder) updateItemById(c *gin.Context) {

}
func (h *Hanlder) deleteItemById(c *gin.Context) {

}
