package migrations

import (
	"database/sql"
	"embed"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
)

//go:embed *.sql
var embedMigrations embed.FS //

func Migrate(db *sql.DB) {
	slog.Info("starting database migrations")
	
	goose.SetDialect("sqlite3")
	goose.SetBaseFS(embedMigrations)

	if err := goose.Up(db, "."); err != nil {
		slog.Error("failed to run migrations", "error", err)
		panic(err)
	}
	
	slog.Info("database migrations completed successfully")
}
