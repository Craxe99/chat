package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

// Константы хранят названия таблиц в базе данных
const (
	usersTable     = "users"
	chatsTable     = "chats"
	usersChatTable = "chat_list"
	messagesTable  = "messages"
)

// Структура конфигурации подключения к базе данных
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

// Функция создания объекта базы данных
func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode))

	if err != nil {
		return db, err
	}

	return db, nil
}
