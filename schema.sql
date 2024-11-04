DROP TABLE IF EXISTS source, post, info;

DROP SEQUENCE IF EXISTS id_seq;

CREATE SEQUENCE id_seq;

-- Схема БД
-- создание таблиц
-- Источники новостей
CREATE TABLE source (
    id INTEGER PRIMARY KEY DEFAULT nextval ('id_seq'),
    url TEXT UNIQUE NOT NULL,
    name TEXT,
    description TEXT
);

-- Новости
CREATE TABLE post (
    id INTEGER PRIMARY KEY DEFAULT nextval ('id_seq'),
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    pub_time INTEGER NOT NULL,
    link TEXT NOT NULL,
    source_id INTEGER NOT NULL,
    FOREIGN KEY (source_id) REFERENCES source (id),
    EXTERNAL_ID TEXT UNIQUE
);

-- Ошибки
CREATE TABLE info (
    id INTEGER PRIMARY KEY DEFAULT nextval ('id_seq'),
    message TEXT NOT NULL,
    time INTEGER NOT NULL,
    type TEXT NOT NULL DEFAULT 'ERROR'
);