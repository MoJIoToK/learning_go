package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"orderApp/pkg/db"
	"orderApp/pkg/model"
	"strconv"
)

type API struct {
	r  *mux.Router //маршрутизатор запросов
	db *db.DB      //БД
}

// NewAPI - конструктор API
func NewAPI(db *db.DB) *API {
	api := API{}
	api.db = db
	api.r = mux.NewRouter()
	api.endpoints()
	return &api
}

// endpoints - функция позволяет зарегестрировать методы API в маршрутизаторе запросов
func (api *API) endpoints() {
	api.r.Use(api.HeadersMiddleware)
	api.r.HandleFunc("/orders", api.getOrdersHandler).Methods(http.MethodGet)
	api.r.HandleFunc("/orders", api.addOrdersHandler).Methods(http.MethodPost)
	api.r.HandleFunc("/orders/{id}", api.updateOrderHandler).Methods(http.MethodPatch)
	api.r.HandleFunc("/orders/{id}", api.deleteOrderHandler).Methods(http.MethodDelete)
}

// Router - метод возвращает маршрутизатор запросов
func (api *API) Router() *mux.Router {
	return api.r
}

// getOrdersHandler - метод возвращает все заказы из БД
func (api *API) getOrdersHandler(w http.ResponseWriter, r *http.Request) {
	// Получение данных из БД
	orders := api.db.GetOrders()
	//Отправка клиенту данных в формате JSON
	json.NewEncoder(w).Encode(orders)
}

// addOrdersHandler - метод добавляет новый заказ в БД
func (api *API) addOrdersHandler(w http.ResponseWriter, r *http.Request) {
	var order model.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id := api.db.AddOrder(order)
	w.Write([]byte(strconv.Itoa(id)))
	w.WriteHeader(http.StatusOK)
}

func (api *API) updateOrderHandler(w http.ResponseWriter, r *http.Request) {
	//Считывание параметра {id} из пути запроса.
	//Например, /orders/45.
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Декодирование в переменную тела запроса, которое должно содержать JSON-представление
	// обновляемого объекта.
	var order model.Order
	err = json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	order.ID = id

	//Обновление данных в БД
	api.db.UpdateOrder(order)

	w.Write([]byte(strconv.Itoa(id)))
	// Отправка клиенту статуса успешного выполнения запроса
	w.WriteHeader(http.StatusOK)
}

func (api *API) deleteOrderHandler(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	api.db.DeleteOrder(id)
	w.WriteHeader(http.StatusOK)
}
