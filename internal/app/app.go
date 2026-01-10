package app

import (
	"flag"
	"log/slog"

	"github.com/msterhuj/ratioarr/internal/router"
	"github.com/msterhuj/ratioarr/internal/config"
	"github.com/msterhuj/ratioarr/internal/crawler"
)

var (
	cfg *config.Config
)

func Run() error {
	configPath := flag.String("config", "config.toml", "Path to config file")
	flag.Parse()
	var err error
	cfg, err = config.Load(*configPath)
	if err != nil {
		slog.Error("failed to load config", "error", err)
		return err
	}
	slog.Info("config loaded successfully")

	crawler.Start()

	r := router.NewRouter()
	r.Run()

	// TODO: init db
	return nil
}
