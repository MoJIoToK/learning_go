package db

import (
	"orderApp/pkg/model"
	"sync"
)

type DB struct {
	m     sync.Mutex          //мьютекс для синхронизации доступа
	id    int                 //текущее значение ID для нового заказа
	store map[int]model.Order //БД заказов
}

func NewDB() *DB {
	db := DB{
		id:    1, //первый номер заказа
		store: map[int]model.Order{},
	}
	return &db
}

// GetOrders - метод возвращает все заказы
func (db *DB) GetOrders() []model.Order {
	db.m.Lock()
	defer db.m.Unlock()

	var order []model.Order
	for _, v := range db.store {
		order = append(order, v)
	}
	return order
}

// AddOrder - метод добавляет новый заказ
func (db *DB) AddOrder(order model.Order) int {
	db.m.Lock()
	defer db.m.Unlock()
	order.ID = db.id
	db.store[db.id] = order
	db.id++
	return order.ID
}

// UpdateOrder - метод обновляет данные заказа по ID
func (db *DB) UpdateOrder(order model.Order) {
	db.m.Lock()
	defer db.m.Unlock()
	db.store[order.ID] = order
}

// DeleteOrder - метод удаляет заказ по ID
func (db *DB) DeleteOrder(id int) {
	db.m.Lock()
	defer db.m.Unlock()
	delete(db.store, id)
}
