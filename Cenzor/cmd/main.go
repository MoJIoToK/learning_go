package main

import (
	"Cenzor/pkg/api"
	"Cenzor/pkg/config"
	"Cenzor/pkg/logger"
	"log/slog"
	"net/http"
)

type server struct {
	api *api.API
	cfg *config.Config
}

func main() {

	logger.SetupLogger()
	slog.Debug("Logger setup load successful")

	cfg := config.MustLoad("./config/config.yaml")
	slog.Debug("Load config file success")

	var s server

	s.cfg = cfg

	s.api = api.New(s.cfg)

	err := http.ListenAndServe(cfg.Address, s.api.Router())
	if err != nil {
		slog.Error("Failed to start server", err)
	}
}
