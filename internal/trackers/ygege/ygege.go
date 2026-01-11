package ygege

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/msterhuj/ratioarr/internal/trackers"
)

var trackerName = "YGEGE"

type Config struct {
	Name   string
	URL    string
	APIKey string
}

type Tracker struct {
	cfg    Config
	client *http.Client
}

func New(cfg Config) *Tracker {
	return &Tracker{
		cfg:    cfg,
		client: &http.Client{},
	}
}

func (t *Tracker) Name() string {
	return t.cfg.Name
}

func (t *Tracker) Type() string {
	return trackerName
}

type YgegeUserResponse struct {
	Username   string  `json:"username"`
	Uploaded   int64   `json:"uploaded"`
	Downloaded int64   `json:"downloaded"`
	Ratio      float64 `json:"ratio"`
}

func (t *Tracker) FetchRatio() (*trackers.Ratio, error) {
	resp, err := t.client.Get(t.cfg.URL + "/user")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}
	// Parse JSON response
	var userResp YgegeUserResponse
	err = json.NewDecoder(resp.Body).Decode(&userResp)
	if err != nil {
		return nil, err
	}

	return &trackers.Ratio{
		Uploaded:   userResp.Uploaded,
		Downloaded: userResp.Downloaded,
		Ratio:      userResp.Ratio,
	}, nil
}

func init() {
	slog.Info("Registering tracker type", "type", trackerName)
	trackers.Register(trackerName, func(cfg any) (trackers.Tracker, error) {
		c, ok := cfg.(Config)
		if !ok {
			return nil, fmt.Errorf("invalid %s config", trackerName)
		}
		return New(c), nil
	})
}
