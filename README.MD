# Решение тестового задания на позицию стажера-бекендера в Avito

Цель задания – разработать чат-сервер, предоставляющий HTTP API для работы с чатами и сообщениями пользователя.

## Установка и запуск программы

Откройте нужный каталог и скачайте в него проект с помощью инструмента git:

```bash
git clone https://github.com/Craxe99/chat
```
Перейдите в каталог со скачанным проектом. Создайте и запустите docker-образ с помощью команды

```bash
docker-compose up --build
```
## Описание программы

Программа написана на языке Go. В качестве хранилища используется база данных PostgreSQL. Добавленные данные
сохраняются после перезапуска сервера благодаря зависимости с папкой ".database", которая автоматически создается при сборке образа.
Сервер приложения запускается на порте 9000 и использует фреймворк GIN. База данных запускается на порте 5436.

Основная архитектура программы разделена на 3 части: обработчик, сервис, репозиторий. Каждая часть
выполняет свои функции и может отдавать команды лишь нижестоящему компоненту архитектуры. Таким образом, репозиторий
не может вызывать методы сервиса, а сервис не может вызывать методы обработчика. Основные методы каждой
части задекларированы в интерфейсах и выполняются через экземпляры объектов, реализующих эти интерфейсы. Данный подход
позволяет добиться гибкости при реализации методов и дает возможность легко заменять отдельные части программы
без надобности в изменении всего кода. Например, можно легко сменить базу данных, изменив реализацию методов репозитория,
при этом обработчик и сервис не будут затронуты.

* Обработчик получает запросы пользователей, обрабатывает их, вызывая методы сервиса, и возвращает результат.
* Сервис получает команды от обработчика и осуществляет основную логику запроса - вычисления, подготовку данных и т.д.
Он может работать с интерфейсами репозитория, вызывая различные методы.
* Репозиторий осуществляет работу с базой данных, получая команды от сервиса.

### Список использованных библиотек
* github.com/gin-gonic/gin - фреймворк GIN
* github.com/jmoiron/sqlx - библиотека для работы с БД
* github.com/lib/pq - драйвер для БД
* github.com/sirupsen/logrus - логгер
* github.com/spf13/viper - инструмент для работы с конфигом

## Основные сущности

Ниже перечислены основные сущности, которыми оперирует сервер.

### User

Пользователь приложения. Имеет следующие свойства:

* **id** - уникальный идентификатор пользователя
* **username** - уникальное имя пользователя
* **created_at** - время создания пользователя

### Chat

Отдельный чат. Имеет следующие свойства:

* **id** - уникальный идентификатор чата
* **name** - уникальное имя чата
* **users** - список пользователей в чате, отношение многие-ко-многим
* **created_at** - время создания

### Message

Сообщение в чате. Имеет следующие свойства:

* **id** - уникальный идентификатор сообщения
* **chat** - ссылка на идентификатор чата, в который было отправлено сообщение
* **author** - ссылка на идентификатор отправителя сообщения, отношение многие-к-одному
* **text** - текст отправленного сообщения
* **created_at** - время создания

## API методы

Методы обрабатывают HTTP запросы c телом, содержащим все необходимые параметры в JSON.

### Добавить нового пользователя

Создание пользователя с уникальным username. Для каждого нового пользователя генерируется личный id.

Входные данные: имя пользователя `username`.

Запрос:

```bash
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"username": "user_1"}' \
  http://localhost:9000/users/add
```

Ответ: `id` созданного пользователя или HTTP-код ошибки.

### Получить список пользователей

Получение списка всех пользователей, отсортированного по дате создания пользователя (от позднего к раннему).

Запрос:

```bash
curl --request GET \
  http://localhost:9000/users/get
```

Ответ: список `users[]`, содержащий поля `id, username, created_at`, или HTTP-код ошибки.

### Создать новый чат между пользователями

Создание чата с указанным названием между выбранными пользователями. Количество пользователей в чате не ограничено,
т. е. можно создать как беседу из множества участников, так и диалог с самим собой. Нельзя продублировать одинакового
пользователя, как нового участника чата.

Входные данные: название чата `name`, массив `users` с `id` пользователей.

Запрос:

```bash
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name": "chat_1", "users": ["<USER_ID_1>", "<USER_ID_2>"]}' \
  http://localhost:9000/chats/add
```

Ответ: `id` созданного чата или HTTP-код ошибки.

### Отправить сообщение в чат от лица пользователя

Отправление сообщения в чат от лица пользователя, который обязан состоять в указанном чате. В противном случае
сообщение не будет отправлено.

Входные данные: id чата `chat`, id пользователя `author`, текст сообщения `text`.

Запрос:

```bash
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"chat": "<CHAT_ID>", "author": "<USER_ID>", "text": "hi"}' \
  http://localhost:9000/messages/add
```

Ответ: `id` созданного сообщения или HTTP-код ошибки.

### Получить список чатов конкретного пользователя

Получение списка чатов, в которых состоит указанный пользователь. Список отсортирован по времени последнего
действия в чате (от позднего к раннему). Новые чаты, в которые еще не написали ни одного сообщения, не могут быть 
отсортированы по дате последнего сообщения, поэтому в качестве параметра времени последнего действия взято время
создания чата.

Входные данные: id пользователя `user`

Запрос:

```bash
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"user": "<USER_ID>"}' \
  http://localhost:9000/chats/get
```

Ответ: список всех чатов `chats[]`, содержащий поля `id, name, users[], created_at, last_action_time`. В списке `users[]`
содержатся данные пользователя `id, username, created_at`.
Или HTTP-код ошибки.

### Получить список сообщений в конкретном чате

Получение списка собщений в указанном чате, отсортированного по дате создания сообщений (от раннего к позднему).

Входные данные: id чата `chat`

Запрос:

```bash
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"chat": "<CHAT_ID>"}' \
  http://localhost:9000/messages/get
```

Ответ: список всех сообщений чата `messages[]`, содержащий поля `id, chat, author, text, created_at`.
Или HTTP-код ошибки.

