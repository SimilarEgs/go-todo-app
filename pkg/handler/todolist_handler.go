package handler

import (
	"net/http"

	"github.com/SimilarEgs/CRUD-TODO-LIST/internal/entity"
	"github.com/gin-gonic/gin"
)

func (h *Hanlder) createList(c *gin.Context) {

	// fetching user ID
	id, ok := c.Get(userCTX)
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
	id, err := h.services.TodoList.CreateList(id.(int64), input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "[Error] operation failed, try again")
	}

	// if operation was successfully done, send code 201 to the client and json with id of the created list
	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})

}
func (h *Hanlder) getListById(c *gin.Context) {
}
func (h *Hanlder) getAllLists(c *gin.Context) {
}
func (h *Hanlder) updateListById(c *gin.Context) {
}
func (h *Hanlder) deleteListById(c *gin.Context) {
}
