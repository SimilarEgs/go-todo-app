package handler

import (
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
		log.Println(err)
		newErrorResponse(c, http.StatusBadRequest, "[Error] invalid request, try again")
		return
	}

	itemId, err := h.services.TodoItem.CreateItem(userId.(int64), int64(listId), input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "[Error] operation failed, try again")
		log.Println(err)
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": itemId,
	})

}
func (h *Hanlder) getItemById(c *gin.Context) {

}
func (h *Hanlder) getAllItems(c *gin.Context) {

}
func (h *Hanlder) updateItemById(c *gin.Context) {

}
func (h *Hanlder) deleteItemById(c *gin.Context) {

}
