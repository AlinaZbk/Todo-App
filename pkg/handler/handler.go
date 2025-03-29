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
		auth.POST("/sign-up", h.signUp) //обработчик регистрации
		auth.POST("/sign-in", h.signIn) //обработчик входа
	}

	api := router.Group("/api")
	{ 
		lists := api.Group("/lists") //будет использована для работы со списком
		{
			lists.POST("/", h.createList) //создание нового списка
			lists.GET("/", h.getAllLists) //получение всех списков
			lists.GET("/:id", h.getListById) //получение списка по id
			lists.PUT("/:id", h.updateList) //редактирование списка по id
			lists.DELETE("/:id", h.deleteList) //удаление списка по id

			items := lists.Group(":id/items") //группа для задач списка
			{
				items.POST("/", h.createItem) //создание новой задачи в списке
				items.GET("/", h.getAllItems) //получение всех задач списка
				items.GET("/:item_id", h.getItemById) //получение задачи по id
				items.PUT("/:items_id", h.updateItem) //редактирование задачи по id
				items.DELETE("/:item_id", h.deleteItem) //удаление задачи по id
			}
		}
	}
	return router
}