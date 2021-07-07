package handler

import (
	"github.com/Craxe99/chat/pkg/service"
	"github.com/gin-gonic/gin"
)

// Обработчики получают запросы клиента по установленным адресам.
// Они отдают команды сервисам, которые осуществляют внутреннюю логику сервера.

// Структура обработчика, хранящая экземпляр сервиса
type Handler struct {
	services *service.Service
}

// Создание нового обработчика
func NewHandler(serv *service.Service) *Handler {
	return &Handler{
		services: serv,
	}
}

// Инициализация зависимостей
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	users := router.Group("/users")
	{
		users.POST("/add", h.addUser)
		users.GET("/get", h.getUsers)
	}

	chats := router.Group("/chats")
	{
		chats.POST("/add", h.addChat)
		chats.POST("/get", h.getChats)
	}

	messages := router.Group("/messages")
	{
		messages.POST("/add", h.addMessage)
		messages.POST("/get", h.getMessages)
	}

	return router
}
