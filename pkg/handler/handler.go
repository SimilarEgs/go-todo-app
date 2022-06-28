package handler

import (
	"github.com/gin-gonic/gin"
)

type Hanlder struct {
}

// initializing project endpoints
func (h *Hanlder) InitRoutes() *gin.Engine {

	// creating router instance
	router := gin.New()

	// creating group of endpoints for authorization and registration
	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
	}

	// creating API endpoints to work with list

	api := router.Group("/api")
	{
		// creating group of lists CRUD
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateListById)
			lists.DELETE("/:id", h.deleteListById)

			// create subsidiary group for lists work
			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET("/:item_id", h.getItemById)
				items.PUT("/:item_id", h.updateItemById)
				items.DELETE("/:item_id", h.deleteItemById)
			}

		}
	}

	return router
}
