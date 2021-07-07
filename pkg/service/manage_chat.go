package service

import (
	"github.com/Craxe99/chat/entities"
	"github.com/Craxe99/chat/pkg/repository"
	"github.com/google/uuid"
	"time"
)

// Структура, реализующая интерфейс управления чатами.
// Хранит экземпляр соответствующего интерфейса репозитория.
type ChatService struct {
	repos repository.ManageChat
}

// Функция создания нового экземпляра структуры ChatService
func newManageChat(repos repository.ManageChat) *ChatService {
	return &ChatService{
		repos: repos,
	}
}

// Функция создания чата
func (c *ChatService) CreateChat(chat entities.Chat) (string, error) {
	// Сохранение времени создания в формате UTC
	chat.CreatedAt = time.Now().UTC()
	// Создание UUID чата с префиксом C (Chat).
	chat.Id = "C" + uuid.NewString()

	// Вызов соответствующей функции интерфейса репозитория и возврат полученных значений
	return c.repos.CreateChat(chat)
}

// Функция получения списка чатов
func (c *ChatService) GetChats(userId string) ([]entities.Chat, error) {
	// Вызов соответствующей функции интерфейса репозитория
	chats, err := c.repos.GetChats(userId)

	// Смена формата времени всех переменных с UTC на Local
	for i, _ := range chats {
		chats[i].CreatedAt = chats[i].CreatedAt.Local()
		chats[i].LastAction = chats[i].LastAction.Local()

		for k, _ := range chats[i].Users {
			chats[i].Users[k].CreatedAt = chats[i].Users[k].CreatedAt.Local()
		}
	}

	// Возврат полученного списка
	return chats, err
}
