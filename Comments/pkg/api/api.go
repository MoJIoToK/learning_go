package api

import (
	"Comments/pkg/model"
	"Comments/pkg/storage"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

// API - программный интерфейс сервиса GoNews
type API struct {
	r       *mux.Router
	storage storage.DB
}

// New - конструктор API.
func New(storage storage.DB) *API {
	api := API{
		storage: storage,
		//Создание роутера
		r: mux.NewRouter(),
	}
	//Запись обработчиков в структуре API
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
	api.r.HandleFunc("/comments/new", api.Add).Methods(http.MethodPost, http.MethodOptions)
	api.r.HandleFunc("/comments/{id:[0-9]+}}", api.Comments).Methods(http.MethodGet, http.MethodOptions)
}

// PostsHandler - метод возвращает n публикации. Где n задаётся пользователем.
func (api *API) Add(w http.ResponseWriter, r *http.Request) {
	const operation = "server.AddComment"

	ct := r.Header.Get("Content-Type")
	media := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))
	if media != "application/json" {
		http.Error(w, "Content-Type header is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	var comment model.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, "cannot decode request", http.StatusBadRequest)
		return
	}

	//if len([]rune(comment.Content)) > ln {
	//	http.Error(w, "the length of the comment must not exceed 1000 characters", http.StatusBadRequest)
	//	return
	//}
	err = api.storage.AddComment(comment)
	//err = api.storage.AddComment(comment)
	if err != nil {
		http.Error(w, "cannot add the comment", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (api *API) Comments(w http.ResponseWriter, r *http.Request) {
	const operation = "server.Comments"

	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "empty post id", http.StatusBadRequest)
		return
	}

	comms, err := api.storage.Comments(id)
	if err != nil {
		http.Error(w, "cannot receive comments", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	enc.SetIndent("", "\t")
	err = enc.Encode(comms)
	if err != nil {
		http.Error(w, "cannot encode comments", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
}
