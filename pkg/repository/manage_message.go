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
	// Проверка, находится ли пользователь в чате
	query := fmt.Sprintf("SELECT id FROM %s WHERE user_id = $1 AND chat_id = $2", usersChatTable)
	var id int
	row := m.db.QueryRow(query, msg.Author, msg.Chat)
	if err := row.Scan(&id); err != nil {
		return "", err
	}

	// Запрос на добавление в таблицу нового сообщения
	query = fmt.Sprintf("INSERT INTO %s (uuid, chat, author, text, created_at) VALUES ($1, $2, $3, $4, $5)", messagesTable)
	_, err := m.db.Exec(query, msg.Id, msg.Chat, msg.Author, msg.Text, msg.CreatedAt)
	if err != nil {
		return "", err
	}

	// Возврат UUID сообщения
	return msg.Id, nil
}

// Функция получения списка сообщений в конкретном чате
func (m *MessagePostgres) GetMessages(chatId string) ([]entities.Message, error) {
	// Объявление среза структуры сообщения
	var messages []entities.Message

	// Запрос на получение сообщений, отправленных в указанный чат, отсортированных по дате создания
	query := fmt.Sprintf("SELECT uuid AS id, chat, author, text, created_at FROM %s WHERE chat = $1 ORDER BY created_at;", messagesTable)
	if err := m.db.Select(&messages, query, chatId); err != nil {
		return nil, err
	} else if messages == nil {
		return nil, nil
	}

	return messages, nil
}
