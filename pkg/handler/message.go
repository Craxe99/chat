package handler

import (
	"github.com/Craxe99/chat/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Функция добавления сообщения
func (h *Handler) addMessage(c *gin.Context) {
	// Входные данные должны соответствовать структуре сообщения
	var input entities.Message

	// Чтение JSON запроса
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Вызов функции создания сообщения через интерфейс сервиса, отвечающий за управление сообщениями
	id, err := h.services.ManageMessage.CreateMessage(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Ответ на запрос
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

// Функция получения списка сообщений в конкретном чате
func (h *Handler) getMessages(c *gin.Context) {
	// Создание временной структуры для получения id чата
	var Input struct {
		Id string `json:"chat" binding:"required"`
	}

	// Чтение JSON запроса
	if err := c.BindJSON(&Input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Вызов функции создания сообщения через интерфейс сервиса, отвечающий за управление сообщениями
	messages, err := h.services.ManageMessage.GetMessages(Input.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if messages == nil {
		newErrorResponse(c, http.StatusBadRequest, "no data")
		return
	}

	// Ответ на запрос
	c.JSON(http.StatusOK, gin.H{
		"messages": messages,
	})
}
