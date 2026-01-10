package crawler

import (
	"context"
	"log/slog"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/msterhuj/ratioarr/internal/repository"
	"github.com/msterhuj/ratioarr/internal/trackers"
)

func Start(trackers []trackers.Tracker, query *repository.Queries) {
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
			err = query.InsertTrackerStat(
				context.Background(),
				repository.InsertTrackerStatParams{
					Name:       tracker.Name(),
					Type:       tracker.Type(),
					Uploaded:   ratio.Uploaded,
					Downloaded: ratio.Downloaded,
					Ratio:      ratio.Ratio,
				},
			)
			if err != nil {
				slog.Error("Failed to insert tracker stat", "tracker", tracker.Name(), "error", err)
			}
		}
	})

	s.StartAsync()
}
