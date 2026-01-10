-- +goose Up
CREATE TABLE IF NOT EXISTS tracker_stats (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
    name TEXT NOT NULL,
    type TEXT NOT NULL,
    uploaded INTEGER NOT NULL,
    downloaded INTEGER NOT NULL,
    ratio REAL NOT NULL
);
-- +goose Down
DROP TABLE IF EXISTS tracker_stats;