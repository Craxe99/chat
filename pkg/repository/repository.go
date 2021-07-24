package repository

import (
	"github.com/Craxe99/chat/entities"
	"github.com/jmoiron/sqlx"
)

// Репозиторий предназначен для работы с базой данных
// Его интерфейсы декларируют требуемые команды.

// Интерфейс управления пользователями
type ManageUser interface {
	CreateUser(user entities.User) (string, error)
	GetUsers() ([]entities.User, error)
}

// Интерфейс управления чатами
type ManageChat interface {
	CreateChat(chat entities.Chat) (string, error)
	GetChats(userId string) ([]entities.Chat, error)
	IsUserInChat(userId string, chatId string) (bool, error)
}

// Интерфейс управления сообщениями
type ManageMessage interface {
	CreateMessage(message entities.Message) (string, error)
	GetMessages(chatId string) ([]entities.Message, error)
}

// Структура репозитория, хранящая объекты, реализующие нужные интерфейсы
type Repository struct {
	ManageUser
	ManageChat
	ManageMessage
}

// Функция создания экземпляра репозитория
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		ManageUser:    newManageUser(db),
		ManageChat:    newManageChat(db),
		ManageMessage: newManageMessage(db),
	}
}
