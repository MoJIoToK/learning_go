package api

import (
	"APIGateWay/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type API struct {
	r *mux.Router
}

func New() *API {
	api := API{
		r: mux.NewRouter(),
	}
	api.endpoints()
	return &api
}

func (api *API) Router() *mux.Router {
	return api.r
}

func (api *API) endpoints() {
	api.r.HandleFunc("/news/latest", api.Latest).Methods(http.MethodGet, http.MethodOptions)
	api.r.HandleFunc("/news/filter", api.Filter).Methods(http.MethodGet, http.MethodOptions)
	api.r.HandleFunc("/news/{id:[0-9]+}", api.Detailed).Methods(http.MethodGet, http.MethodOptions)
	api.r.HandleFunc("/news/commet", api.AddComment).Methods(http.MethodPost, http.MethodOptions)
}

func (api *API) Latest(w http.ResponseWriter, r *http.Request) {

	param := r.URL.Query().Get("page")
	if param == "" {
		param = "1"
	}

	page, err := strconv.Atoi(param)
	if err != nil {
		http.Error(w, "incorrect page number", http.StatusBadRequest)
		return
	}
	_ = page

	var resp []model.NewsShortDetailed
	w.Header().Set("Access-Control-Allow-Origin", "*")
	resp = model.HardCode
	enc := json.NewEncoder(w)
	err = enc.Encode(&resp)
	if err != nil {
		http.Error(w, "failed to encode news", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
}

func (api *API) Filter(w http.ResponseWriter, r *http.Request) {
	var resp []model.NewsShortDetailed
	w.Header().Set("Access-Control-Allow-Origin", "*")
	resp = model.HardCode
	enc := json.NewEncoder(w)
	err := enc.Encode(resp)
	if err != nil {
		http.Error(w, "failed to encode news", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
}

func (api *API) Detailed(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	var resp = model.NewsFullDetailed{
		News:     model.HardCode[id-1],
		Comments: model.CommentNews1,
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	err = enc.Encode(&resp)
	if err != nil {
		http.Error(w, "failed to encode detailed news", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
}

func (api *API) AddComment(w http.ResponseWriter, r *http.Request) {
	var req []model.Comment
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "failed to decode news", http.StatusInternalServerError)
		return
	}
	if req[0].Content == "" {
		http.Error(w, "Empty Comment", http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusCreated)
}
