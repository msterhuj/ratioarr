package migrations

import (
	"database/sql"
	"embed"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
)

//go:embed *.sql
var embedMigrations embed.FS //

func Migrate(db *sql.DB) {
	goose.SetDialect("sqlite3")
	goose.SetBaseFS(embedMigrations)

	if err := goose.Up(db, "."); err != nil {
		panic(err)
	}
}
