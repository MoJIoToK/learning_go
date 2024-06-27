/*
-- Создание таблиц
DROP TABLE IF EXISTS users, tasks, labels, tasks_labels;
-- Создание таблицы users
CREATE TABLE users(
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL
);
-- Создание таблицы tasks
CREATE TABLE tasks(
	id SERIAL PRIMARY KEY,
	opened BIGINT NOT NULL DEFAULT extract(epoch from now()), 
	closed BIGINT NOT NULL DEFAULT 0,
	author_id INTEGER REFERENCES users(id) DEFAULT 0,
	assigned_id INTEGER REFERENCES users(id) DEFAULT 0,
	title TEXT,
	content TEXT	
);
-- Создание таблицы labels
CREATE TABLE labels(
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL
);
-- Создание таблицы tasks_labels
CREATE TABLE tasks_labels(
	task_id INTEGER REFERENCES tasks(id),
	label_id INTEGER REFERENCES labels(id)
);
-- Наполнение БД начальными данными
INSERT INTO users (id, name) VALUES(0, 'Nick');

INSERT INTO tasks(title, content) VALUES 
	('Hello', 'Hello msg'),
	('GOODBYE', 'BYE BYE');
*/
--SELECT * FROM tasks;
/*
INSERT INTO labels (id , name) VALUES 
	(0, 'Done'),
	(1, 'Ready');

INSERT INTO tasks_labels (task_id , label_id) VALUES 
	(1, 0),
	(2, 1);
*/
--SELECT * FROM tasks_labels;

SELECT * FROM tasks, tasks_labels WHERE tasks_labels.label_id = 1;