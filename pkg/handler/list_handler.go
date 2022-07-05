package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Hanlder) createList(c *gin.Context) {

	id, _ := c.Get(userCTX)
	c.JSON(http.StatusOK, map[string]interface{}{
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
