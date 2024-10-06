package main

import (
	"log"
	"module31_pratice/pkg/api"
	"module31_pratice/pkg/storage"
	"module31_pratice/pkg/storage/postgres"
	"net/http"
)

// server - структура сервера GoNews
type server struct {
	db  storage.Interface
	api *api.API
}

func main() {
	var srv server

	// Создание объектов баз данных
	// БД на основе слайсов
	//db1 := memdb.New()

	//БД на основе postgresql
	conn := "postgres://postgres:password@localhost:5432/posts"
	db2, err := postgres.New(conn)
	if err != nil {
		log.Fatal(err)
	}

	//БД на основе mongoDB
	//conn = "mongodb://localhost:27017/"
	//db3, err := mongo.New(conn)
	//if err != nil {
	//	log.Fatal(err)
	//}

	_ = db2

	// Инициализация хранилище сервера конкретной БД
	srv.db = db2

	// Создание объекта API
	srv.api = api.New(srv.db)

	//Запуск сервера на порту 8080.
	http.ListenAndServe(":80", srv.api.Router())

}
