package entities

import "time"

// Структура пользователя
type User struct {
	Id        string    `json:"id" db:"id"`
	Username  string    `json:"username" db:"username" binding:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// Структура чата
type Chat struct {
	Id         string    `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Users      []User    `json:"users" db:"users"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	LastAction time.Time `json:"last_action_time" db:"last_action"`
}

// Структура сообщения
type Message struct {
	Id        string    `json:"id" db:"id"`
	Chat      string    `json:"chat" db:"chat" binding:"required"`
	Author    string    `json:"author" db:"author" binding:"required"`
	Text      string    `json:"text" db:"text" binding:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
