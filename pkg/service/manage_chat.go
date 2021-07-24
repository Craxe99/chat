package service

import (
	"github.com/Craxe99/chat/entities"
	"github.com/Craxe99/chat/pkg/repository"
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

// Функция проверки, находится ли пользователь в чате
func (c *ChatService) IsUserInChat(userId string, chatId string) (bool, error) {
	return c.repos.IsUserInChat(userId, chatId)
}
