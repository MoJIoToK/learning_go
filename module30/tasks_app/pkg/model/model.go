package model

// Task модель задачи
type Task struct {
	ID          int
	Opened      int64
	Closed      int64
	Author_id   int
	Assigned_id int
	Title       string
	Content     string
}
