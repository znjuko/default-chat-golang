DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS chats CASCADE;
DROP TABLE IF EXISTS messages;
DROP TABLE IF EXISTS chat_user CASCADE;
DROP TABLE IF EXISTS emoji;
DROP TABLE IF EXISTS online CASCADE;
DROP TABLE IF EXISTS newmessages CASCADE;

CREATE EXTENSION IF NOT EXISTS CITEXT;


CREATE TABLE users
(
    u_id     SERIAL PRIMARY KEY ,
    login    CITEXT UNIQUE,
    password TEXT
);

CREATE TABLE chats
(
    ch_id        SERIAL PRIMARY KEY ,
    name         TEXT,
    last_msg_id INT DEFAULT 0,
    last_msg_log CITEXT,
    last_msg_txt TEXT DEFAULT ''
);

CREATE TABLE chat_user
(
    cu_id SERIAL,
    ch_id INT NOT NULL REFERENCES chats (ch_id),
    u_id  INT NOT NULL REFERENCES users (u_id)
);

CREATE UNIQUE INDEX idx_un_ch ON chat_user(ch_id,u_id);

CREATE TABLE emoji
(
    em_id     SERIAL PRIMARY KEY ,
    main_word TEXT,
    slug      TEXT
);

CREATE TABLE messages
(
    msg_id SERIAL PRIMARY KEY ,
    u_id INT NOT NULL REFERENCES users(u_id),
    ch_id  INT NOT NULL REFERENCES chats (ch_id),
    text   TEXT
);

CREATE TABLE newmessages
(
    msg_id INT REFERENCES messages,
    u_id INT REFERENCES users
);

CREATE TABLE online
(
    onl_id SERIAL,
    u_id INT UNIQUE REFERENCES users (u_id)
);