package api

import (
	"GoNews/internal/storage"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type API struct {
	r       *mux.Router
	storage storage.DB
}

// New - конструктор API
func New(storage storage.DB) *API {
	api := API{
		storage: storage,
		r:       mux.NewRouter(),
	}
	api.endpoints()
	return &api
}

// Router - метод возвращает маршрутизатор для использования в качестве
// аргумента HTTP-сервера.
func (api *API) Router() *mux.Router {
	return api.r
}

// endpoints - метод регистрирует методы API в маршрутизаторе запросов.
func (api *API) endpoints() {
	api.r.HandleFunc("/news/{id:[0-9]+}", api.PostsHandler).Methods(http.MethodGet, http.MethodOptions)
	api.r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./cmd/webapp"))))
}

// PostsHandler -
func (api *API) PostsHandler(w http.ResponseWriter, r *http.Request) {
	const operation = "API.PostsHandler"

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == http.MethodOptions {
		return
	}

	vars := mux.Vars(r)
	n, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("%s: Failed to convert id to int", operation)
	}

	news, err := api.storage.GetPosts(n)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("%s: Failed to get posts from DB", operation)
		return
	}

	json.NewEncoder(w).Encode(news)
}
