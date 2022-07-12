package handler

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, msg string) {
	// logging error
	log.Error(msg)
	// if error -> block and prevent pending handlers -> send as resp http code and msg body
	c.AbortWithStatusJSON(statusCode, errorResponse{msg})
}
