package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Hanlder) createList(c *gin.Context) {

}
func (h *Hanlder) getListById(c *gin.Context) {

	c.String(http.StatusOK, "Hello world")
}
func (h *Hanlder) getAllLists(c *gin.Context) {
}
func (h *Hanlder) updateListById(c *gin.Context) {

}
func (h *Hanlder) deleteListById(c *gin.Context) {

}
