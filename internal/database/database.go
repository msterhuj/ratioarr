package database

import (
	"database/sql"
	"log/slog"
)

func Connect() (*sql.DB, error) {
	slog.Info("initializing database")
	db, err := sql.Open("sqlite3", "ratioarr.db")
	if err != nil {
		slog.Error("failed to open database", "error", err)
		return nil, err
	}
	slog.Info("database initialized successfully")
	return db, nil
}
