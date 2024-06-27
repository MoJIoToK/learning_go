package main

import (
	"fmt"
	"log"
	"tasks_app/pkg/model"
	"tasks_app/pkg/storage"
)

func main() {
	var err error
	conn := "postgres://postgres:password@localhost:5432/tasks"

	db, err := storage.New(conn)
	if err != nil {
		log.Fatal(err)
	}

	//Проверка содержимого БД
	tasks, err := db.GetTasks()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("tasks before operation:", tasks)

	//Добавление новой записи
	//newTask := model.Task{
	//	Author_id: 0,
	//	Title:     "New Task",
	//	Content:   "This is new task",
	//}
	//taskId, err := db.NewTask(newTask)
	//if err != nil {
	//	log.Println(err)
	//}
	//fmt.Println("New task:", taskId)

	//Проверка содержимого БД после добавления новой записи
	tasks1, err := db.GetTasks()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("tasks after add new task:", tasks1)

	//Поиск записи по ID
	taskById, err := db.GetTaskById(1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("task by id:", taskById)

	//Поиск записей по ID автора
	tasksByAuthorId, err := db.GetTaskByAuthorId(0)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("tasks by author Id:", tasksByAuthorId)

	//Поиск записей по лейблу
	tasksByLabelId, err := db.GetTaskByLabelId("Ready")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("tasks by label Id:", tasksByLabelId)

	//Обновление записи по ID записи
	taskByUpdate := model.Task{
		Author_id: 1,
		Title:     "Update Task",
		Content:   "This is update task",
	}

	if err = db.UpdateTaskById(3, taskByUpdate); err != nil {
		log.Println(err)
	}

	tasksAfterUpdate, err := db.GetTasks()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("tasks after update:", tasksAfterUpdate)

	//Удаление записи по ID
	if err = db.DeleteTaskById(3); err != nil {
		log.Println(err)
	}

	tasksAfterDelete, err := db.GetTasks()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("tasks after delete:", tasksAfterDelete)
}
