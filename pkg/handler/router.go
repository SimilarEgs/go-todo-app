package handler

import (
	"github.com/SimilarEgs/CRUD-TODO-LIST/pkg/service"
	"github.com/gin-gonic/gin"
)

type Hanlder struct {
	services *service.Service // our handlers will envoke service methods
}

// implementing service dependencies
func NewHandler(services *service.Service) *Hanlder {
	return &Hanlder{services: services}
}

// initializing project endpoints
func (h *Hanlder) InitRoutes() *gin.Engine {

	router := gin.New()

	// creating group of endpoints for authorization and registration
	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
	}

	// creating API endpoints of lists
	// and passing  middlewere authentication handler
	api := router.Group("/api", h.userAuthMiddleware)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateListById)
			lists.DELETE("/:id", h.deleteListById)

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
