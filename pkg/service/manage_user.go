package service

import (
	"github.com/Craxe99/chat/entities"
	"github.com/Craxe99/chat/pkg/repository"
	"github.com/google/uuid"
	"time"
)

// Структура, реализующая интерфейс управления пользователями.
// Хранит экземпляр соответствующего интерфейса репозитория.
type UserService struct {
	repos repository.ManageUser
}

// Функция создания нового экземпляра структуры UserService
func newManageUser(repos repository.ManageUser) *UserService {
	return &UserService{
		repos: repos,
	}
}

// Функция создания пользователя
func (u *UserService) CreateUser(user entities.User) (string, error) {
	// Сохранение времени создания в формате UTC
	user.CreatedAt = time.Now().UTC()
	// Создание UUID пользователя с префиксом U (User).
	user.Id = "U" + uuid.NewString()

	// Вызов соответствующей функции интерфейса репозитория и возврат полученных значений
	return u.repos.CreateUser(user)
}

// Функция получения списка пользователей
func (u *UserService) GetUsers() ([]entities.User, error) {
	// Вызов соответствующей функции интерфейса репозитория
	users, err := u.repos.GetUsers()

	// Смена формата времени создания пользователя с UTC на Local
	for i, _ := range users {
		users[i].CreatedAt = users[i].CreatedAt.Local()
	}

	// Возврат полученного списка
	return users, err
}
