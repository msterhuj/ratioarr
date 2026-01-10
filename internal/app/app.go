package app

import (
	"flag"
	"fmt"

	"github.com/msterhuj/ratioarr/internal/config"
)

func Run() error {
	configPath := flag.String("config", "config.toml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}
	fmt.Println("Config loaded:", cfg)

	fmt.Println("Starting application with config:", *configPath)

	// TODO: init logger
	// TODO: init db
	// TODO: start crawler
	// TODO: start http server
	return nil
}
