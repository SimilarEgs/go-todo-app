package handler

import (
	"net/http"

	"github.com/SimilarEgs/CRUD-TODO-LIST/internal/entity"
	"github.com/gin-gonic/gin"
)

func (h *Hanlder) signUp(c *gin.Context) {
	// var to store user data
	var input entity.User

	// validate request body
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// affter parsing and data validation, send data to the service layer via «CreateUser» method
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// if operation was successfully done, send code 201 to the user and json with id of created user
	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})

}
func (h *Hanlder) signIn(c *gin.Context) {

}
