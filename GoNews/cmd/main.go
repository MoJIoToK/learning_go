package main

import (
	"GoNews/internal/api"
	"GoNews/internal/rss"
	"GoNews/internal/storage"
	"GoNews/internal/storage/memdb"
	"log"
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

	posts, err := rss.Parse()
	if err != nil {
		log.Fatal(err)
	}

	for _, post := range posts {
		id, err := srv.db.AddPost(post)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Added post #%d to the database", id)
	}

	http.ListenAndServe(":80", srv.api.Router())

}
