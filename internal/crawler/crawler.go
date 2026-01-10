package crawler

import (
	"log/slog"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/msterhuj/ratioarr/internal/trackers"
)

func Start(trackers []trackers.Tracker) {
	s := gocron.NewScheduler(time.UTC)

	s.Every(100).Seconds().Do(func() {
		slog.Info("Scheduled task executed", "time", time.Now())
		for _, tracker := range trackers {
			ratio, err := tracker.FetchRatio()
			if err != nil {
				slog.Error("Failed to fetch ratio", "tracker", tracker.Name(), "error", err)
				continue
			}
			slog.Info("Fetched ratio", "tracker", tracker.Name(), "uploaded", ratio.Uploaded, "downloaded", ratio.Downloaded, "ratio", ratio.Ratio)
		}
	})

	s.StartAsync()
}
