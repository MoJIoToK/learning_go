package main

import (
	"log"
	"net/http"
	"orderApp/pkg/model"
	"time"

	"orderApp/pkg/api"
	"orderApp/pkg/db"
)

func main() {
	// Инициализация БД в памяти.
	db := db.NewDB()

	p := []model.Product{
		{
			Name:  "Apple",
			Price: 1000,
		},
		{
			Name:  "Orange",
			Price: 2000,
		},
	}
	order := model.Order{
		IsOpen:       true,
		DeliveryTime: time.Now().Unix(),
		Products:     p,
	}
	db.AddOrder(order)

	// Создание объекта API, использующего БД в памяти.
	api := api.NewAPI(db)
	// Запуск сетевой службы и HTTP-сервера
	// на всех локальных IP-адресах на порту 80.
	err := http.ListenAndServe(":80", api.Router())
	if err != nil {
		log.Fatal(err)
	}
}
