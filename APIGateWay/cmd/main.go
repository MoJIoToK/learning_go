package main

import (
	"APIGateWay/pkg/api"
	"APIGateWay/pkg/config"
	"APIGateWay/pkg/logger"
	"log"
	"log/slog"
	"net/http"
)

type server struct {
	api *api.API
}

func main() {

	logger.SetupLogger()
	slog.Debug("Logger setup load successful")

	cfg := config.MustLoad("./config/config.yaml")
	slog.Debug("Load config file success")

	var s server

	s.api = api.New(cfg)

	err := http.ListenAndServe(":80", s.api.Router())
	if err != nil {
		log.Fatal(err)
	}

}
