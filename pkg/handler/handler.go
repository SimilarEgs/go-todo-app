package handler

import (
	"github.com/gin-gonic/gin"
)

type Hanlder struct {
}

// initializing project endpoints
func (h *Hanlder) initRoutes() *gin.Engine {

	// creating router instance
	router := gin.New()

	// creating group of endpoints for authorization and registration
	auth := router.Group("/auth")
	{
		auth.POST("/sign-in")
		auth.POST("/sign-up")
	}

	// creating API endpoints to work with list

	api := router.Group("/api")
	{
		// creating lists CRUD group
		lists := api.Group("/lists")
		{
			lists.POST("/")
			lists.GET("/id")
			lists.GET("/id:")
			lists.PUT("/id:")
			lists.DELETE("/id:")

			// create subsidiary group for lists work
			items := lists.Group(":id/items")
			{
				items.POST("/")
				items.GET("/")
				items.GET("/:item_id")
				items.PUT("/:item_id")
				items.DELETE("/:item_id:")
			}

		}
	}

	return router
}
