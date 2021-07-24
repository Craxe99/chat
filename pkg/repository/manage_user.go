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
	// Объявление переменной, хранящей id созданного пользователя
	var id string
	// Запрос на добавление в таблицу нового пользователя
	query := fmt.Sprintf("INSERT INTO %s (username, created_at) VALUES ($1, $2) RETURNING id", usersTable)
	// Выполнение запроса с переданными параметрами
	row := u.db.QueryRow(query, user.Username, user.CreatedAt)

	if err := row.Scan(&id); err != nil {
		return "", err
	}

	// Возврат ID пользователя
	return id, nil
}

func (u *UserPostgres) GetUsers() ([]entities.User, error) {
	// Объявление среза структуры пользователя
	var users []entities.User

	// Запрос на получение списка пользователей, отсортированного по дате создания по убыванию
	query := fmt.Sprintf("SELECT id, username, created_at FROM %s ORDER BY created_at DESC;", usersTable)
	if err := u.db.Select(&users, query); err != nil {
		return nil, err
	} else if users == nil {
		return nil, nil
	}

	// Возврат среза структуры пользователей
	return users, nil
}
