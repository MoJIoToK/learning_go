package storage

import "testing/mock/pkg/model"

// Интерфейс БД.
// Этот интерфейс позволяет абстрагироваться от конкретной СУБД.
// Можно создать реализацию БД в памяти для модульных тестов.
type Interface interface {
	Tasks() ([]model.Task, error)
	NewTask(model.Task) (int, error)
}
