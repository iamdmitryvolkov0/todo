package handler

import (
	"github.com/gin-gonic/gin"
	"todo/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/signup", h.signUp)
		auth.POST("/signin", h.signIn)
	}
	api := router.Group("/api", h.userIdentify)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.allLists)
			lists.GET("/:id", h.listById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group("/:id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.allItems)
			}
		}
		items := api.Group("/items")
		{
			items.GET("/:id", h.item)
			items.PUT("/:id", h.updateItem)
			items.DELETE("/:id", h.deleteItem)
		}
	}
	return router
}
