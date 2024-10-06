package main

import (
	"APIGateWay/pkg/api"
	"log"
	"net/http"
)

type server struct {
	api *api.API
}

func main() {

	var s server

	s.api = api.New()

	err := http.ListenAndServe(":80", s.api.Router())
	if err != nil {
		log.Fatal(err)
	}

}
