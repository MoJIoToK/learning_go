package api

import (
	"GoNews/internal/storage"
	"github.com/gorilla/mux"
	"net/http"
)

type API struct {
	r       *mux.Router
	storage storage.DB
}

func New(storage storage.DB) *API {
	api := API{
		storage: storage,
		r:       mux.NewRouter(),
	}
	api.endpoints()
	return &api
}

func (api *API) endpoints() {
	api.r.HandleFunc("/posts", api.PostsHandler).Methods(http.MethodGet, http.MethodOptions)
}

func (api *API) Router() *mux.Router {
	return api.r
}

func (api *API) PostsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//news, err := api.storage.GetPosts()
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//bytes, err := json.Marshal(news)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	w.Write([]byte("Hello"))
}
