package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Conn(path string) (*sql.DB, error) {
	dsn := "Side-B.db?_foreign_keys=on&_journal_mode=WAL"
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return nil, pingErr
	}

	return db, nil
}
