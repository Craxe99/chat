package repository

import (
	"fmt"
	"github.com/Craxe99/chat/entities"
	"github.com/jmoiron/sqlx"
)

// Структура, реализующая интерфейс управления чатами.
// Хранит ссылку на базу данных.
type ChatPostgres struct {
	db *sqlx.DB
}

// Функция создания нового экземпляра структуры ChatPostgres
func newManageChat(db *sqlx.DB) *ChatPostgres {
	return &ChatPostgres{
		db: db,
	}
}

// Функция создания чата
func (c *ChatPostgres) CreateChat(chat entities.Chat) (string, error) {
	// Начало транзакции
	tx, err := c.db.Begin()
	if err != nil {
		return "", err
	}

	// Добавление записи в таблицу чатов
	query := fmt.Sprintf("INSERT INTO %s (uuid, name, created_at) VALUES ($1, $2, $3)", chatsTable)
	_, err = tx.Exec(query, chat.Id, chat.Name, chat.CreatedAt)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	// Добавление всех пользователей чата в таблицу, связывающую таблицы чата и пользователей по принципу многие ко многим
	for _, user := range chat.Users {
		query = fmt.Sprintf("INSERT INTO %s (user_id, chat_id) VALUES ($1, $2)", usersChatTable)
		_, err = tx.Exec(query, user.Id, chat.Id)
		if err != nil {
			tx.Rollback()
			return "", err
		}
	}

	// Подтверждение транзакции и возврат UUID чата
	return chat.Id, tx.Commit()
}

// Функция получения списка чатов
func (c *ChatPostgres) GetChats(userId string) ([]entities.Chat, error) {
	// Объявление среза структуры чата
	var chats []entities.Chat

	// Запрос на получение всех полей чатов, в которых находится пользователь с требуемым ID.
	// Подтаблица t1 в качестве последнего действия (last_action) отмечает время последнего отправленного в чат сообщение.
	// Таблицы из которых она состоит объединены по принципу left join, начиная с таблицы сообщений.
	// Таким образом, можно с уверенностью сказать, что в рассматриваемом чате полученной подтаблицы есть хотя бы одно сообщение.
	// Если чат был создан, но в него не было отправлено ни одного сообщения, то он не отобразится в этой подтаблице.
	// Чтобы не потерять такие чаты, подтаблица t1 объединяется с подтаблицей t2, в которой те же таблицы объединяются по
	// принципу right outer join в том же порядке. Это позволяет найти чаты, в которых нет сообщений.
	// В качетсве параметра для сортировки last_action в подтаблице t2 используется время создания чата.
	// Полученные подтаблицы объединяются и сортируются.
	// Благодаря этому чаты без сообщений правильно сортируются совместно с остальными чатами, сохраняя временную хронологию.
	query := fmt.Sprintf(
		`select id, name, created_at, last_action
		from (
		select a.chat as id, c.name, c.created_at, max(a.created_at) as last_action
		from %s a left join %s b on a.chat = b.chat_id
		left join %s c on b.chat_id = c.uuid
		where b.user_id = $1
		group by a.chat, c.name, c.created_at) t1
		union all
		select id, name, created_at, last_action
		from (
		select c.uuid as id, c.name, c.created_at, max(c.created_at) as last_action
		from %s a right join %s b on a.chat = b.chat_id
		right join %s c on b.chat_id = c.uuid
		where b.user_id = $1 and a.created_at is null
		group by c.uuid, c.name, c.created_at) t2
		order by last_action desc`,
		messagesTable, usersChatTable, chatsTable, messagesTable, usersChatTable, chatsTable)
	if err := c.db.Select(&chats, query, userId); err != nil {
		return nil, err
	} else if chats == nil {
		return nil, nil
	}

	// Для каждого элемента полученного среза чатов выполняется запрос на получение списка пользователей, которые находятся в конкретном чате.
	for i, _ := range chats {
		query = fmt.Sprintf("SELECT uuid as id, username, created_at FROM %s a join %s b on a.user_id = b.uuid WHERE chat_id = $1", usersChatTable, usersTable)
		if err := c.db.Select(&chats[i].Users, query, chats[i].Id); err != nil {
			return nil, err
		}
	}

	// Возврат среза чатов
	return chats, nil
}
