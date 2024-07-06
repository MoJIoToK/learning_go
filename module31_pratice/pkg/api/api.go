package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"module31_pratice/pkg/model"
	"module31_pratice/pkg/storage"
	"net/http"
	"strconv"
)

// API - программный интерфейс сервиса GoNews
type API struct {
	db     storage.Interface
	router *mux.Router
}

// New - конструктор для структуры API
func New(db storage.Interface) *API {
	api := API{
		db: db,
	}
	//Создание роутера
	api.router = mux.NewRouter()

	//Запись обработчиков структуре API
	api.endpoints()

	return &api
}

// endpoints - регистрация обработчиков
func (api *API) endpoints() {
	api.router.HandleFunc("/posts", api.postsHandler).Methods(http.MethodGet, http.MethodOptions)
	api.router.HandleFunc("/posts", api.addPostHandler).Methods(http.MethodPost, http.MethodOptions)
	api.router.HandleFunc("/posts/update/{id:[0-9]+}", api.updatePostHandler).Methods(http.MethodPut, http.MethodOptions)
	api.router.HandleFunc("/posts/delete/{id:[0-9]+}", api.deletePostHandler).Methods(http.MethodDelete, http.MethodOptions)
}

// Router - возвращает маршрутизатор запросов. Необходим для передачи маршрутизатора веб-серверу.
func (api *API) Router() *mux.Router {
	return api.router
}

// postHandler - метод возвращает все публикации
func (api *API) postsHandler(w http.ResponseWriter, r *http.Request) {

	//Вызов метода получения записей из БД
	posts, err := api.db.GetPosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Сериализация данных в JSON
	bytes, err := json.Marshal(posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}

// addPostHandler - метод позволяет добавить публикацию
func (api *API) addPostHandler(w http.ResponseWriter, r *http.Request) {
	var post model.Post

	//Десериализация JSON формата в структуру model.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Вызов метода добавления публикации в БД
	id, err := api.db.AddPost(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := fmt.Sprintf("Публикация %v добавлена", id)

	w.Write([]byte(response))
	w.WriteHeader(http.StatusOK)
}

// updatePostHandler - метод позволяет обновить публикацию
func (api *API) updatePostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var post model.Post
	//Десериализация JSON формата в структуру model.Post
	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Вызов метода обновления публикации в БД
	err = api.db.UpdatePost(id, post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := fmt.Sprintf("Публикация %v изменена", id)

	w.Write([]byte(response))
	w.WriteHeader(http.StatusOK)
}

// deletePostHandler - метод позволяет удалить публикацию
func (api *API) deletePostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Вызов метода удаления публикации в БД
	err = api.db.DeletePost(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := fmt.Sprintf("Публикация %v удалена", id)

	w.Write([]byte(response))
	w.WriteHeader(http.StatusOK)
}
