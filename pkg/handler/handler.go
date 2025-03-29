package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New() //создаем новый роутер

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up") //обработчик регистрации
		auth.POST("/sign-in") //обработчик входа
	}

	api := router.Group("/api")
	{ 
		lists := api.Group("/lists") //будет использована для работы со списком
		{
			lists.POST("/") //создание нового списка
			lists.GET("/") //получение всех списков
			lists.GET("/:id") //получение списка по id
			lists.PUT("/:id") //редактирование списка по id
			lists.DELETE("/:id") //удаление списка по id

			items := lists.Group(":id/items") //группа для задач списка
			{
				items.POST("/") //создание новой задачи в списке
				items.GET("/") //получение всех задач списка
				items.GET("/:item_id") //получение задачи по id
				items.PUT("/:items_id") //редактирование задачи по id
				items.DELETE("/:item_id") //удаление задачи по id
			}
		}
	}
	return router
}