package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Структура с текстом ошибки
type errorResponse struct {
	Message string `json:"message"`
}

// Функция отменяющая запрос с выводом ошибки
func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
