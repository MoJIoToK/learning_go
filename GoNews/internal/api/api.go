package api

import (
	"GoNews/internal/storage"
	"github.com/gorilla/mux"
	"net/http"
)

type API struct {
	r       *mux.Router
	storage *storage.Storage
}

func New(storage *storage.Storage) *API {
	api := API{}
	api.storage = storage
	api.r = mux.NewRouter()
	api.endpoints()
	return &api
}

func (api *API) endpoints() {

}

func (api *API) Router() *mux.Router {
	return api.r
}

func (api *API) getPostsHandler(w http.ResponseWriter, r *http.Request) {

}
