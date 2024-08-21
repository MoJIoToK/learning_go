package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"orderApp/pkg/db"
	"orderApp/pkg/model"
	"strconv"
	"testing"
)

var (
	//Создаём чистый объект API для теста
	dbase = db.NewDB()
	api   = NewAPI(dbase)
)

func TestAPI_addOrdersHandler(t *testing.T) {

	order := model.Order{}

	orderJSON, err := json.Marshal(order)
	if err != nil {
		t.Fatalf("Не удалось закодировать запрос к серверу: %v", err)
	}

	//Создаём HTTP-запрос
	req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewReader(orderJSON))

	//Создаём объект для записи ответа обработчика
	rr := httptest.NewRecorder()

	//Вызываем маршрутизатор. Маршрутизатор для пути и метода запроса вызовет обработчик.
	//Обработчик запишет ответ в созданный объект.
	api.r.ServeHTTP(rr, req)
	if !(rr.Code == http.StatusOK) {
		t.Errorf("Код неверен: получен %d, а необходим %d", rr.Code, http.StatusOK)
	}

	//Чтение тела ответа
	b, err := io.ReadAll(rr.Body)
	if err != nil {
		t.Fatalf("Не удалось раскодировать ответ сервера: %v", err)
	}

	//Парсинг возвращенного ID заказа из тела ответа
	id, err := strconv.Atoi(string(b))

	//Раскодирование JSON в массив заказов
	if id != 1 {
		t.Fatalf("Не удалось раскодировать ответ сервера: %v", err)
	}
}

func TestAPI_getOrdersHandler(t *testing.T) {

	dbase.AddOrder(model.Order{})

	//Создаём HTTP-запрос
	req := httptest.NewRequest(http.MethodGet, "/orders", nil)

	//Создаём объект для записи ответа обработчика
	rr := httptest.NewRecorder()

	//Вызываем маршрутизатор. Маршрутизатор для пути и метода запроса вызовет обработчик.
	//Обработчик запишет ответ в созданный объект.
	api.r.ServeHTTP(rr, req)
	if !(rr.Code == http.StatusOK) {
		t.Errorf("Код неверен: получен %d, а необходим %d", rr.Code, http.StatusOK)
	}

	//Чтение тела ответа
	b, err := io.ReadAll(rr.Body)
	if err != nil {
		t.Fatalf("Не удалось раскодировать ответ сервера: %v", err)
	}

	//Раскодирование JSON в массив заказов
	var data []model.Order
	err = json.Unmarshal(b, &data)
	if err != nil {
		t.Fatalf("Не удалось раскодировать ответ сервера: %v", err)
	}

	//Проверка наличия в массиве ровно одного элемента
	const wantLen = 1
	if len(data) != wantLen {
		t.Fatalf("Получено %d записей, ожидалось %d", len(data), wantLen)
	}
}

func TestAPI_updateOrderHandler(t *testing.T) {

	order := model.Order{}

	dbase.AddOrder(order)

	updatedOrder := model.Order{}
	orderJSON, err := json.Marshal(updatedOrder)
	if err != nil {
		t.Fatalf("Не удалось закодировать запрос к серверу: %v", err)
	}

	//Создаём HTTP-запрос
	req := httptest.NewRequest(http.MethodPatch, "/orders/1", bytes.NewReader(orderJSON))

	//Создаём объект для записи ответа обработчика
	rr := httptest.NewRecorder()

	//Вызываем маршрутизатор. Маршрутизатор для пути и метода запроса вызовет обработчик.
	//Обработчик запишет ответ в созданный объект.
	api.r.ServeHTTP(rr, req)
	if !(rr.Code == http.StatusOK) {
		t.Errorf("Код неверен: получен %d, а необходим %d", rr.Code, http.StatusOK)
	}
}

func TestAPI_deleteOrderHandler(t *testing.T) {

	order := model.Order{}

	dbase.AddOrder(order)

	//Создаём HTTP-запрос
	req := httptest.NewRequest(http.MethodDelete, "/orders/1", nil)

	//Создаём объект для записи ответа обработчика
	rr := httptest.NewRecorder()

	//Вызываем маршрутизатор. Маршрутизатор для пути и метода запроса вызовет обработчик.
	//Обработчик запишет ответ в созданный объект.
	api.r.ServeHTTP(rr, req)
	if !(rr.Code == http.StatusOK) {
		t.Errorf("Код неверен: получен %d, а необходим %d", rr.Code, http.StatusOK)
	}
}
