package storage

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"tasks_app/pkg/model"
)

// Storage - структура для БД
type Storage struct {
	db *pgxpool.Pool
}

// New - конструктор для БД.
func New(constr string) (*Storage, error) {
	db, err := pgxpool.New(context.Background(), constr)
	if err != nil {
		return nil, err
	}
	return &Storage{db: db}, nil
}

// NewTask - метод сохраняет в БД задачу.
func (s *Storage) NewTask(task model.Task) (int, error) {
	var id int
	query := "INSERT INTO tasks (title, content, author_id) VALUES ($1, $2, $3) RETURNING id"

	err := s.db.QueryRow(context.Background(), query, task.Title, task.Content, task.Author_id).Scan(&id)
	return id, err
}

// GetTasks - метод позволяющий получить все задачи из БД.
func (s *Storage) GetTasks() ([]model.Task, error) {
	query := "SELECT * FROM tasks"

	rows, err := s.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	var tasks []model.Task

	for rows.Next() {
		var task model.Task
		err = rows.Scan(
			&task.ID,
			&task.Opened,
			&task.Closed,
			&task.Author_id,
			&task.Assigned_id,
			&task.Title,
			&task.Content,
		)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, rows.Err()
}

// GetTaskById - метод возвращает задачу по её ID.
func (s *Storage) GetTaskById(id int) (model.Task, error) {
	query := "SELECT * FROM tasks WHERE id = $1"

	var task model.Task
	err := s.db.QueryRow(context.Background(), query, id).
		Scan(
			&task.ID,
			&task.Opened,
			&task.Closed,
			&task.Author_id,
			&task.Assigned_id,
			&task.Title,
			&task.Content,
		)
	return task, err
}

func (s *Storage) GetTaskByAuthorId(id int) ([]model.Task, error) {
	query := "SELECT * FROM tasks WHERE author_id = $1"

	rows, err := s.db.Query(context.Background(), query, id)
	if err != nil {
		return nil, err
	}

	var tasks []model.Task

	for rows.Next() {
		var task model.Task
		err = rows.Scan(
			&task.ID,
			&task.Opened,
			&task.Closed,
			&task.Author_id,
			&task.Assigned_id,
			&task.Title,
			&task.Content,
		)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, rows.Err()
}

// GetTaskByLabelId - метод возвращающий список задач по пометке
func (s *Storage) GetTaskByLabelId(label string) ([]model.Task, error) {
	query := "SELECT tasks.* FROM tasks " +
		"JOIN tasks_labels ON tasks.id = tasks_labels.task_id " +
		"JOIN labels ON tasks_labels.label_id = labels.id " +
		"WHERE labels.name = $1"

	rows, err := s.db.Query(context.Background(), query, label)
	if err != nil {
		return nil, err
	}

	var tasks []model.Task

	for rows.Next() {
		var task model.Task
		err = rows.Scan(
			&task.ID,
			&task.Opened,
			&task.Closed,
			&task.Author_id,
			&task.Assigned_id,
			&task.Title,
			&task.Content,
		)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, rows.Err()
}

// UpdateTaskById - метод обновляет запись по ID
func (s *Storage) UpdateTaskById(id int, task model.Task) error {
	query := "UPDATE tasks SET title = $1, content=$2, author_id=$3 WHERE id = $4 RETURNING id"
	_, err := s.db.Exec(context.Background(), query, task.Title, task.Content, task.Author_id, id)
	return err
}

// DeleteTaskById - метод удаляет задачу
func (s *Storage) DeleteTaskById(id int) error {
	query := "DELETE FROM tasks WHERE id = $1"
	_, err := s.db.Exec(context.Background(), query, id)
	return err
}
