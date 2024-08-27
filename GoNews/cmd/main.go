package main

import (
	"GoNews/internal/api"
	"GoNews/internal/config"
	"GoNews/internal/model"
	"GoNews/internal/rss"
	"GoNews/internal/storage"
	"GoNews/internal/storage/mongo"
	"log"
	"net/http"
	"time"
)

type server struct {
	db  storage.DB
	api *api.API
}

func main() {
	var srv server

	cfg := config.MustLoad("./config/config.yaml")

	//db := memdb.New()

	conn := cfg.StoragePath
	db, err := mongo.New(conn)
	if err != nil {
		log.Fatal(err)
	}

	srv.db = db
	srv.api = api.New(srv.db)

	chPosts := make(chan []model.Post)
	chErrors := make(chan error)

	for _, url := range cfg.URLS {
		go parseURL(url, chPosts, chErrors, cfg.Period)
	}

	go func() {
		for posts := range chPosts {
			db.AddPost(posts)
		}
	}()

	go func() {
		for err := range chErrors {
			log.Println("Error:", err)
		}
	}()

	err = http.ListenAndServe(":80", srv.api.Router())
	if err != nil {
		log.Fatal(err)
	}

}

func parseURL(url string, posts chan []model.Post, errs chan error, period int) {
	for {
		news, err := rss.Parse(url)
		if err != nil {
			errs <- err
			continue
		}
		posts <- news
	}
	time.Sleep(time.Second * time.Duration(period))
}
