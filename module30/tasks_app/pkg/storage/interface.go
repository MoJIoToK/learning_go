package storage

import "tasks_app/pkg/model"

type Interface interface {
	NewTask(model.Task) (int, error)
	GetTasks() ([]model.Task, error)
	GetTaskById(int) (model.Task, error)
	GetTaskByAuthorId(int) ([]model.Task, error)
	GetTaskByLabelId(string) ([]model.Task, error)
	UpdateTaskById(int, model.Task) error
	DeleteTaskById(int) error
}
