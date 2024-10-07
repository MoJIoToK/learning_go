package main

import (
	"Comments/pkg/api"
	"Comments/pkg/storage"
	"Comments/pkg/storage/mongo"
	"log"
	"net/http"
)

// server - структура сервера, хранящая экземпляр БД и api.
type server struct {
	db  storage.DB
	api *api.API
}

func main() {

	var srv server

	//cfg := config.MustLoad("./config/config.yaml")

	//Инициализация зависимостей
	conn := "mongodb://localhost:27017/"
	db, err := mongo.New(conn)
	if err != nil {
		log.Fatal(err)
	}

	srv.db = db
	srv.api = api.New(srv.db)

	//Запуск веб-сервера с API и приложением
	err = http.ListenAndServe(":80", srv.api.Router())
	if err != nil {
		log.Fatal(err)
	}
}
