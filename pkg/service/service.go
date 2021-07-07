package service

import (
	"github.com/Craxe99/chat/entities"
	"github.com/Craxe99/chat/pkg/repository"
)

// Сервис предназначен для реализации внутренней логики сервера.
// Его интерфейсы декларируют требуемые команды.
// Он может отдавать команды репозиторию с помощью интерфейсов репозитория.

// Интерфейс управления пользователями
type ManageUser interface {
	CreateUser(user entities.User) (string, error)
	GetUsers() ([]entities.User, error)
}

// Интерфейс управления чатами
type ManageChat interface {
	CreateChat(chat entities.Chat) (string, error)
	GetChats(userId string) ([]entities.Chat, error)
}

// Интерфейс управления сообщениями
type ManageMessage interface {
	CreateMessage(message entities.Message) (string, error)
	GetMessages(chatId string) ([]entities.Message, error)
}

// Структура сервиса, хранящая объекты, реализующие нужные интерфейсы
type Service struct {
	ManageUser
	ManageChat
	ManageMessage
}

// Функция создания экземпляра сервиса
func NewService(repos *repository.Repository) *Service {
	return &Service{
		ManageUser:    newManageUser(repos.ManageUser),
		ManageChat:    newManageChat(repos.ManageChat),
		ManageMessage: newManageMessage(repos.ManageMessage),
	}
}
