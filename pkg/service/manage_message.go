package service

import (
	"github.com/Craxe99/chat/entities"
	"github.com/Craxe99/chat/pkg/repository"
	"time"
)

// Структура, реализующая интерфейс управления сообщениями.
// Хранит экземпляр соответствующего интерфейса репозитория.
type MessageService struct {
	repos repository.ManageMessage
}

// Функция создания нового экземпляра структуры MessageService
func newManageMessage(repos repository.ManageMessage) *MessageService {
	return &MessageService{
		repos: repos,
	}
}

// Функция создания сообщения
func (m *MessageService) CreateMessage(message entities.Message) (string, error) {
	// Сохранение времени создания в формате UTC
	message.CreatedAt = time.Now().UTC()

	// Вызов соответствующей функции интерфейса репозитория и возврат полученных значений
	return m.repos.CreateMessage(message)
}

// Функция получения списка сообщений
func (m *MessageService) GetMessages(chatId string) ([]entities.Message, error) {
	// Вызов соответствующей функции интерфейса репозитория
	messages, err := m.repos.GetMessages(chatId)

	// Смена формата времени создания сообщения с UTC на Local
	for i, _ := range messages {
		messages[i].CreatedAt = messages[i].CreatedAt.Local()
	}

	// Возврат полученного списка
	return messages, err
}
