package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Trackers []TrackerConfig `toml:"trackers"`
}

type TrackerConfig struct {
	Name   string `toml:"name"`
	Type   string `toml:"type"`
	Url    string `toml:"url"`
	ApiKey string `toml:"api_key"`
}

func Load(path string) (*Config, error) {
	if _, err := os.Stat(path); err != nil {
		return nil, fmt.Errorf("config file not found: %w", err)
	}

	var cfg Config
	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		return nil, fmt.Errorf("invalid config file: %w", err)
	}

	return &cfg, nil
}
