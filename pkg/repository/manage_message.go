package repository

import (
	"fmt"
	"github.com/Craxe99/chat/entities"
	"github.com/jmoiron/sqlx"
)

// Структура, реализующая интерфейс управления сообщениями.
// Хранит ссылку на базу данных.
type MessagePostgres struct {
	db *sqlx.DB
}

// Функция создания нового экземпляра структуры MessagePostgres
func newManageMessage(db *sqlx.DB) *MessagePostgres {
	return &MessagePostgres{
		db: db,
	}
}

// Функция создания сообщения
func (m *MessagePostgres) CreateMessage(msg entities.Message) (string, error) {
	// Объявление переменной, хранящей id созданного сообщения
	var id string

	// Запрос на добавление в таблицу нового сообщения
	query := fmt.Sprintf("INSERT INTO %s (chat, author, text, created_at) VALUES ($1, $2, $3, $4) RETURNING id", messagesTable)
	row := m.db.QueryRow(query, msg.Chat, msg.Author, msg.Text, msg.CreatedAt)
	if err := row.Scan(&id); err != nil {
		return "", err
	}

	// Возврат ID сообщения
	return id, nil
}

// Функция получения списка сообщений в конкретном чате
func (m *MessagePostgres) GetMessages(chatId string) ([]entities.Message, error) {
	// Объявление среза структуры сообщения
	var messages []entities.Message

	// Запрос на получение сообщений, отправленных в указанный чат, отсортированных по дате создания
	query := fmt.Sprintf("SELECT id, chat, author, text, created_at FROM %s WHERE chat = $1 ORDER BY created_at;", messagesTable)
	if err := m.db.Select(&messages, query, chatId); err != nil {
		return nil, err
	} else if messages == nil {
		return nil, nil
	}

	return messages, nil
}
