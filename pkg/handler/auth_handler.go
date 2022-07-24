package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/SimilarEgs/go-todo-app/internal/entity"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// @Summary     SignUp
// @Tags        Auth
// @Description Creates new user account
// @ID          create-user
// @Accept      json
// @Produce     json
// @Param       input         body      entity.User true "data for registering a new user"
// @Success     201           {integer} integer			 1
// @Failure     400 	      {object}  errorResponse
// @Failure     404           {object}  errorResponse
// @Failure     500           {object}  errorResponse
// @Failure     default       {object}  errorResponse
// @Router      /auth/sign-up [post]
func (h *Hanlder) signUp(c *gin.Context) {
	// var for storing user input data
	var input entity.User

	// validate request body
	if err := c.BindJSON(&input); err != nil {
		msg := fmt.Sprintf("[Error] invalid request, try again: %v", err)
		newErrorResponse(c, http.StatusBadRequest, msg)
		return
	}

	// affter parsing and data validation, send data to the service layer via «CreateUser» method
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		msg := fmt.Sprintf("[Error] operation failed, try again: %v", err)
		newErrorResponse(c, http.StatusInternalServerError, msg)
		return
	}

	// if operation was successfully done, send code 201 to the client and json with id of created user
	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})

}

// struct for parsing request body
type singInUserInput struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

// @Summary     SignIn
// @Tags        Auth
// @Description login using user credentials
// @ID          login
// @Accept      json
// @Produce     json
// @Param       input   	   body     singInUserInput true "login credentials"
// @Success     200     	   {string} string          "token"
// @Failure     400     	   {object} errorResponse
// @Failure     404     	   {object} errorResponse
// @Failure     500     	   {object} errorResponse
// @Failure     default        {object} errorResponse
// @Router      /auth/sign-in  [post]
func (h *Hanlder) signIn(c *gin.Context) {

	var input singInUserInput

	// validate request body
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "[Error] invalid login credentials")
		return
	}

	// generate JWT and error handling
	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		// error handling for 400 -> unknown users
		if err == sql.ErrNoRows {
			newErrorResponse(c, http.StatusBadRequest, "[Error] accout with given username not found")
			return
		}
		// error handilng for 401 -> incorrect pwd
		if err == bcrypt.ErrMismatchedHashAndPassword {
			newErrorResponse(c, http.StatusUnauthorized, "[Error] invalid login credentials")
			return
		}

		msg := fmt.Sprintf("[Error] connection error, try again: %v", err)
		newErrorResponse(c, http.StatusInternalServerError, msg)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
