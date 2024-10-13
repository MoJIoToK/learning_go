package api

import (
	"Cenzor/pkg/config"
	"Cenzor/pkg/logger"
	"Cenzor/pkg/middleware"
	"Cenzor/pkg/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"strings"
)

// API - программный интерфейс сервиса Cenzor
type API struct {
	r   *mux.Router
	cfg *config.Config
}

// New - конструктор API.
func New(cfg *config.Config) *API {
	api := API{
		r:   mux.NewRouter(),
		cfg: cfg,
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
	api.r.Use(middleware.RequestID)
	api.r.Use(middleware.Logger)
	api.r.HandleFunc("/", api.Cenzor).Methods(http.MethodPost, http.MethodOptions)
}

func (api *API) Cenzor(w http.ResponseWriter, r *http.Request) {

	const operation = "server.Censor"

	log := slog.Default().With(
		slog.String("op", operation),
		slog.String("request_id", middleware.GetRequestID(r.Context())),
	)

	log.Info("Request to censor comment")

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	var req model.Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Error("Cannot decode request", logger.Err(err))
		http.Error(w, "cannot decode request", http.StatusBadRequest)
		return
	}

	if isOffensive(req.Content, api.cfg.CensorList) {
		log.Info("Comment contains offensive words")
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	log.Info("Comment is allowed")

	w.WriteHeader(http.StatusOK)
	log.Info("Request served successfully")
}

// isOffensive - проверяет контент комментария на содержание недопустимых выражений.
func isOffensive(text string, words []string) bool {
	for _, word := range words {
		if strings.Contains(text, word) {
			return true
		}
	}
	return false
}
