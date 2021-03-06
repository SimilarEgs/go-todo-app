package handler

import (
	_ "github.com/SimilarEgs/go-todo-app/docs"
	"github.com/SimilarEgs/go-todo-app/pkg/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Hanlder struct {
	services *service.Service
}

// implementing service dependencies
func NewHandler(services *service.Service) *Hanlder {
	return &Hanlder{services: services}
}

// initializing project endpoints
func (h *Hanlder) InitRoutes() *gin.Engine {

	router := gin.New()

	// initilzing swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// creating group of endpoints for authorization and registration
	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
	}

	// creating API endpoints of todolists
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
			}

		}

		items := api.Group("items")
		{
			items.GET("/:item_id", h.getItemById)
			items.PUT("/:item_id", h.updateItemById)
			items.DELETE("/:item_id", h.deleteItemById)
		}
	}

	return router
}
