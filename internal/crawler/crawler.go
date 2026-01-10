package crawler

import (
	"log/slog"
	"time"

	"github.com/go-co-op/gocron"
)

func Start() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(10).Seconds().Do(func() {
		slog.Info("Scheduled task executed", "time", time.Now())
	})

	s.StartAsync()
}
