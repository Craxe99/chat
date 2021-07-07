CREATE TABLE users
(
    id serial primary key,
    uuid varchar(37) not null unique,
    username varchar(255) not null unique,
    created_at timestamp
);

CREATE TABLE chats
(
    id serial primary key,
    uuid varchar(37) not null unique,
    name varchar(255) not null,
    created_at timestamp
);

CREATE TABLE chat_list
(
    id serial primary key,
    user_id varchar(37) references users(uuid) on delete cascade not null,
    chat_id varchar(37) references chats(uuid) on delete cascade not null,
    constraint no_duplicates unique (user_id, chat_id)
);

CREATE TABLE messages
(
    id serial primary key,
    uuid varchar(37) not null unique,
    chat varchar(37) references chats(uuid) on delete cascade not null,
    author varchar(37) references users(uuid) not null,
    text varchar(4095),
    created_at timestamp
);

