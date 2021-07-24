CREATE TABLE users
(
    id serial primary key,
    username varchar(255) not null unique,
    created_at timestamp
);

CREATE TABLE chats
(
    id serial primary key,
    name varchar(255) not null,
    created_at timestamp
);

CREATE TABLE chat_list
(
    id serial primary key,
    user_id int references users(id) on delete cascade not null,
    chat_id int references chats(id) on delete cascade not null,
    constraint no_duplicates unique (user_id, chat_id)
);

CREATE TABLE messages
(
    id serial primary key,
    chat int references chats(id) on delete cascade not null,
    author int references users(id) not null,
    text varchar(4095),
    created_at timestamp
);



