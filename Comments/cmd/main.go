package main

import (
	"Comments/pkg/api"
	"Comments/pkg/config"
	"Comments/pkg/logger"
	"Comments/pkg/storage"
	"Comments/pkg/storage/mongo"
	"log"
	"log/slog"
	"net/http"
)

// server - структура сервера, хранящая экземпляр БД и api.
type server struct {
	db  storage.DB
	api *api.API
}

func main() {

	logger.SetupLogger()
	slog.Debug("Logger setup load successful")

	var srv server

	cfg := config.MustLoad("./config/config.yaml")
	slog.Debug("Load config file success")

	//Инициализация зависимостей
	db, err := mongo.New(cfg.StoragePath)
	if err != nil {
		log.Fatal(err)
	}

	srv.db = db
	srv.api = api.New(srv.db)

	//Запуск веб-сервера с API и приложением
	err = http.ListenAndServe(cfg.Address, srv.api.Router())
	if err != nil {
		log.Fatal(err)
	}
}
