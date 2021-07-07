package handler

import (
	"github.com/Craxe99/chat/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Функция добавления чата
func (h *Handler) addChat(c *gin.Context) {
	// Создание временной структуры для входных данных
	var Input struct {
		Name  string   `json:"name" binding:"required"`
		Users []string `json:"users" binding:"required"`
	}

	// Объявление экземпляра структуры чата, для передачи в сервис
	var chat entities.Chat

	// Чтение JSON запроса
	if err := c.BindJSON(&Input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	for _, str := range Input.Users {
		chat.Users = append(chat.Users, entities.User{Id: str})
	}
	chat.Name = Input.Name

	// Вызов функции создания чата через интерфейс сервиса, отвечающий за управление чатами
	id, err := h.services.ManageChat.CreateChat(chat)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Ответ на запрос
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

// Функция получения списка чатов конкретного пользователя
func (h *Handler) getChats(c *gin.Context) {
	// Создание временной структуры для получения id пользователя
	var Input struct {
		Id string `json:"user" binding:"required"`
	}

	// Чтение JSON запроса
	if err := c.BindJSON(&Input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Вызов функции получения списка чатов через интерфейс сервиса, отвечающий за управление чатами
	chats, err := h.services.ManageChat.GetChats(Input.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if chats == nil {
		newErrorResponse(c, http.StatusBadRequest, "no data")
		return
	}

	// Ответ на запрос
	c.JSON(http.StatusOK, gin.H{
		"chats": chats,
	})
}
