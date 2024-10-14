package api

import (
	"Comments/pkg/logger"
	"Comments/pkg/middleware"
	"Comments/pkg/model"
	"Comments/pkg/storage"
	"Comments/pkg/tree"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"log/slog"
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
	api.r.Use(middleware.Logger)
	api.r.Use(middleware.RequestID)
	api.r.HandleFunc("/news/comments/new", api.Add).Methods(http.MethodPost, http.MethodOptions)
	api.r.HandleFunc("/news/id/{id}", api.Comments).Methods(http.MethodGet, http.MethodOptions)
}

func (api *API) Add(w http.ResponseWriter, r *http.Request) {
	const operation = "server.AddComment"

	log := slog.Default().With(
		slog.String("op", operation),
		slog.String("request_id", middleware.GetRequestID(r.Context())),
	)

	log.Info("Request to add comment")

	ct := r.Header.Get("Content-Type")
	media := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))
	if media != "application/json" {
		log.Error("Content-Type header is not application/json")
		http.Error(w, "Content-Type header is not application/json", http.StatusUnsupportedMediaType)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	var comment model.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		log.Error("Cannot decode request", logger.Err(err))
		http.Error(w, "cannot decode request", http.StatusBadRequest)
		return
	}
	log.Debug("request body decoded")

	if comment.Content == "" {
		log.Error("Comment has empty content field")
		http.Error(w, "Empty comment", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	id, err := api.storage.AddComment(ctx, comment)
	if err != nil {
		log.Error("Cannot add comment to DB", logger.Err(err))
		if errors.Is(err, storage.ErrIncorrectParentID) || errors.Is(err, storage.ErrIncorrectPostID) {
			http.Error(w, "Incorrect data", http.StatusBadRequest)
			return
		}
		http.Error(w, "cannot add the comment", http.StatusInternalServerError)
		return
	}
	log.Debug("comment added to DB successfully", slog.String("id", id))

	w.WriteHeader(http.StatusCreated)
	log.Info("Request served successfully")
}

func (api *API) Comments(w http.ResponseWriter, r *http.Request) {
	const operation = "server.Comments"

	log := slog.Default().With(
		slog.String("op", operation),
		slog.String("request_id", middleware.GetRequestID(r.Context())),
	)

	log.Info("request to receive comments")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		log.Error("Empty news id")
		http.Error(w, "Empty news id", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	comms, err := api.storage.Comments(ctx, id)
	if err != nil {
		log.Error("Cannot receive comments", logger.Err(err))
		if errors.Is(err, storage.ErrNoComments) {
			http.Error(w, "News id not found", http.StatusNotFound)
			return
		}
		if errors.Is(err, storage.ErrIncorrectPostID) {
			http.Error(w, "incorrect post id", http.StatusBadRequest)
			return
		}
		http.Error(w, "Cannot receive comments", http.StatusInternalServerError)
		return
	}

	log.Debug("Comments received successfully")

	root, err := tree.Build(comms)
	if err != nil {
		log.Error("cannot build comments tree", logger.Err(err))
		http.Error(w, "cannot receive comments", http.StatusInternalServerError)
		return
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "\t")
	err = enc.Encode(root.Comments)
	if err != nil {
		log.Error("Cannot encode comments", logger.Err(err))
		http.Error(w, "Cannot encode comments", http.StatusInternalServerError)
		return
	}

	log.Info("Request served successfully")

}
