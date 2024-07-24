package memdb

import "testing/mock/pkg/model"

type DB []model.Task

// Выполнение контракта интерфейса storage.Interface
func (db *DB) Tasks() ([]model.Task, error) {
	return *db, nil
}
func (db *DB) NewTask(task model.Task) (int, error) {
	*db = append(*db, task)
	return 0, nil
}
