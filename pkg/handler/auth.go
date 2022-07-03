package handler

import (
	"net/http"

	"github.com/SimilarEgs/CRUD-TODO-LIST/internal/entity"
	"github.com/gin-gonic/gin"
)

func (h *Hanlder) signIn(c *gin.Context) {
	// var to store user data
	var input entity.User

	// validate request body
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	
	// affter parsing and data validation, send data to the service layer

}
func (h *Hanlder) signUp(c *gin.Context) {

}
