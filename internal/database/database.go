package database

import (
	"database/sql"
	"log/slog"
	"os"
	"path/filepath"
)

func Connect() (*sql.DB, error) {
	slog.Info("initializing database")

	dbPath := "ratioarr.db"

	// Ensure the directory exists
	dbDir := filepath.Dir(dbPath)
	if dbDir != "." {
		if err := os.MkdirAll(dbDir, 0755); err != nil {
			slog.Error("failed to create database directory", "error", err)
			return nil, err
		}
	}

	// Create the database file if it doesn't exist
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		slog.Info("database file doesn't exist, creating it", "path", dbPath)
		file, err := os.Create(dbPath)
		if err != nil {
			slog.Error("failed to create database file", "error", err)
			return nil, err
		}
		file.Close()
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		slog.Error("failed to open database", "error", err)
		return nil, err
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		slog.Error("failed to ping database", "error", err)
		db.Close()
		return nil, err
	}

	slog.Info("database initialized successfully")
	return db, nil
}
