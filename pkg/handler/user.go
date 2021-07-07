package handler

import (
	"github.com/Craxe99/chat/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Функция добавления пользователя
func (h *Handler) addUser(c *gin.Context) {
	// Входные данные должны соответствовать структуре пользователя
	var input entities.User

	// Чтение JSON запроса
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Вызов функции создания пользователя через интерфейс сервиса, отвечающий за управление пользователями
	id, err := h.services.ManageUser.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Ответ на запрос
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

// Функция получения списка пользователей
func (h *Handler) getUsers(c *gin.Context) {
	// Вызов функции получения списка через интерфейс сервиса, отвечающий за управление пользователями
	users, err := h.services.ManageUser.GetUsers()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if users == nil {
		newErrorResponse(c, http.StatusInternalServerError, "Users count is null")
		return
	}

	// Ответ на запрос
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
