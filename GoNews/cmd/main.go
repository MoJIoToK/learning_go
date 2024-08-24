package main

import (
	"GoNews/internal/api"
	"GoNews/internal/storage"
	"GoNews/internal/storage/memdb"
	"net/http"
)

type server struct {
	db  storage.DB
	api *api.API
}

func main() {
	var srv server

	db := memdb.New()

	srv.db = db

	srv.api = api.New(srv.db)

	http.ListenAndServe(":8080", srv.api.Router())

}
