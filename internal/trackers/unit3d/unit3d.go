package unit3d

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/TheVovchenskiy/data"
	"github.com/msterhuj/ratioarr/internal/trackers"
)

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

type Unit3DUserResponse struct {
	Username   string `json:"username"`
	Uploaded   string `json:"uploaded"`
	Downloaded string `json:"downloaded"`
	Ratio      string `json:"ratio"`
	Seeding    int    `json:"seeding"`
	Leeching   int    `json:"leeching"`
	HitAndRuns int    `json:"hit_and_runs"`
}

func (t *Tracker) FetchRatio() (*trackers.Ratio, error) {
	resp, err := t.client.Get(t.cfg.URL + "/api/user?api_token=" + t.cfg.APIKey)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}
	// Parse JSON response
	var userResp Unit3DUserResponse
	err = json.NewDecoder(resp.Body).Decode(&userResp)
	if err != nil {
		return nil, err
	}

	uploadedStr := strings.ReplaceAll(userResp.Uploaded, " ", "")
	downloadedStr := strings.ReplaceAll(userResp.Downloaded, " ", "")

	// Convert Uploaded/Downloaded strings to bytes
	uploadedBytes, err := data.ParseSize(uploadedStr)
	if err != nil {
		return nil, fmt.Errorf("erreur parsing uploaded: %v", err)
	}

	downloadedBytes, err := data.ParseSize(downloadedStr)
	if err != nil {
		return nil, fmt.Errorf("erreur parsing downloaded: %v", err)
	}

	// Convert Ratio string to float64
	ratio, err := strconv.ParseFloat(userResp.Ratio, 64)
	if err != nil {
		return nil, fmt.Errorf("erreur parsing ratio: %v", err)
	}

	return &trackers.Ratio{
		Uploaded:   uploadedBytes.B(),
		Downloaded: downloadedBytes.B(),
		Ratio:      ratio,
	}, nil
}

func init() {
	trackers.Register("UNIT3D", func(cfg any) (trackers.Tracker, error) {
		c, ok := cfg.(Config)
		if !ok {
			return nil, fmt.Errorf("invalid UNIT3D config")
		}
		return New(c), nil
	})
}
