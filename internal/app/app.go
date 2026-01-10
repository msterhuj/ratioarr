package app

import (
	"flag"
	"log/slog"

	"github.com/msterhuj/ratioarr/internal/config"
	"github.com/msterhuj/ratioarr/internal/crawler"
	"github.com/msterhuj/ratioarr/internal/router"
	"github.com/msterhuj/ratioarr/internal/trackers"
	"github.com/msterhuj/ratioarr/internal/trackers/unit3d"
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
	slog.Info("config", "config", cfg)

	var allTrackers []trackers.Tracker
	for _, tcfg := range cfg.Trackers {
		switch tcfg.Type {
		case "UNIT3D":
			slog.Info("registering UNIT3D tracker", "name", tcfg.Name)
			tracker, err := trackers.New(tcfg.Type, unit3d.Config{
				Name:   tcfg.Name,
				URL:    tcfg.Url,
				APIKey: tcfg.ApiKey,
			})
			if err != nil {
				slog.Error("failed to create UNIT3D tracker", "error", err)
				return err
			}
			allTrackers = append(allTrackers, tracker)
		default:
			slog.Warn("unknown tracker type", "type", tcfg.Type)
		}
	}

	crawler.Start(allTrackers)

	r := router.NewRouter()
	r.Run()

	// TODO: init db
	return nil
}
