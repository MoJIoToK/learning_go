-- Создание таблиц
DROP TABLE IF EXISTS users, tasks, labels, tasks_labels;

-- Создание таблицы users
CREATE TABLE users
(
    id   SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

-- Создание таблицы tasks
CREATE TABLE tasks
(
    id          SERIAL PRIMARY KEY,
    opened      BIGINT NOT NULL               DEFAULT extract(epoch from now()),
    closed      BIGINT NOT NULL               DEFAULT 0,
    author_id   INTEGER REFERENCES users (id) DEFAULT 0,
    assigned_id INTEGER REFERENCES users (id) DEFAULT 0,
    title       TEXT,
    content     TEXT
);

-- Создание таблицы labels
CREATE TABLE labels
(
    id   SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

-- Создание таблицы tasks_labels
CREATE TABLE tasks_labels
(
    task_id  INTEGER REFERENCES tasks (id),
    label_id INTEGER REFERENCES labels (id)
);

-- Наполнение БД начальными данными
INSERT INTO users (id, name)
VALUES (0, 'Nick');

SELECT *
FROM users;