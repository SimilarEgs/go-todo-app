package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey = "Authorization"
	authorizationType      = "Bearer"
	userCTX                = "userId"
)

// this function gets authentication values from headers
// validates this received values
// and passes them to the context
func (h *Hanlder) userAuthMiddleware(c *gin.Context) {

	// getting auth header
	authHeader := c.GetHeader(authorizationHeaderKey)

	// validating 401
	if authHeader == "" {
		newErrorResponse(c, http.StatusUnauthorized, "[Error] unauthorized, authorization token is missing")
		return
	}

	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		newErrorResponse(c, http.StatusUnauthorized, "[Error] unauthorized, invalid authorization header format")
		return
	}

	authType := fields[0]
	if authType != authorizationType {
		err := fmt.Errorf("[Error] unauthorized, unsoported authorization type %s", authType)
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	// parse token
	jwtToken := fields[1]
	userId, err := h.services.Authorization.ParseToken(jwtToken)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "[Error] unauthorized, invalid authorization token")
	}

	// setting user id in ctx for use in further handlers
	c.Set(userCTX, userId)
}
