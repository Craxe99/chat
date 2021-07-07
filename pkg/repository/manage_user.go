package repository

import (
	"fmt"
	"github.com/Craxe99/chat/entities"
	"github.com/jmoiron/sqlx"
)

// Структура, реализующая интерфейс управления пользователями.
// Хранит ссылку на объект базы данных.
type UserPostgres struct {
	db *sqlx.DB
}

// Функция создания нового экземпляра структуры UserPostgres
func newManageUser(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{
		db: db,
	}
}

// Функция создания пользователя
func (u *UserPostgres) CreateUser(user entities.User) (string, error) {
	// Запрос на добавление в таблицу нового пользователя
	query := fmt.Sprintf("INSERT INTO %s (uuid, username, created_at) VALUES ($1, $2, $3)", usersTable)
	// Выполнение запроса с переданными параметрами
	_, err := u.db.Exec(query, user.Id, user.Username, user.CreatedAt)
	if err != nil {
		return "", err
	}

	// Возврат UUID пользователя
	return user.Id, nil
}

func (u *UserPostgres) GetUsers() ([]entities.User, error) {
	// Объявление среза структуры пользователя
	var users []entities.User

	// Запрос на получение списка пользователей, отсортированного по дате создания по убыванию
	query := fmt.Sprintf("SELECT uuid AS id, username, created_at FROM %s ORDER BY created_at DESC;", usersTable)
	if err := u.db.Select(&users, query); err != nil {
		return nil, err
	} else if users == nil {
		return nil, nil
	}

	// Возврат среза структуры пользователей
	return users, nil
}
