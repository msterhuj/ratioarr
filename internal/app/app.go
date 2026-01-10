package app

import (
	"database/sql"
	"flag"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"
	"github.com/msterhuj/ratioarr/internal/config"
	"github.com/msterhuj/ratioarr/internal/crawler"
	"github.com/msterhuj/ratioarr/internal/database"
	"github.com/msterhuj/ratioarr/internal/migrations"
	"github.com/msterhuj/ratioarr/internal/repository"
	"github.com/msterhuj/ratioarr/internal/router"
	"github.com/msterhuj/ratioarr/internal/trackers"
	"github.com/msterhuj/ratioarr/internal/trackers/unit3d"
)

var (
	cfg     *config.Config
	db      *sql.DB
	queries *repository.Queries
)

func Run() error {

	configPath := flag.String("config", "config.toml", "Path to config file")
	disableCrawler := flag.Bool("disable-crawler", false, "Disable the crawler")
	flag.Parse()
	var err error
	cfg, err = config.Load(*configPath)
	if err != nil {
		slog.Error("failed to load config", "error", err)
		return err
	}
	slog.Info("config loaded successfully")
	slog.Info("config", "config", cfg)

	db, err = database.Connect()
	if err != nil {
		slog.Error("failed to connect to database", "error", err)
		return err
	}
	defer db.Close()
	migrations.Migrate(db)
	queries = repository.New(db)

	slog.Info("initializing trackers")
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

	if !*disableCrawler {
		slog.Info("starting ratio crawler")
		crawler.Start(allTrackers, queries)
	} else {
		slog.Warn("ratio crawler is disabled ratio will not be updated")
	}

	r := router.NewRouter()
	r.Run()
	return nil
}
