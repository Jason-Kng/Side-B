package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Creates a connection to the db, also sets up the tables if they dont already exist
func InitDB(filepath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}

	query := `
	CREATE TABLE IF NOT EXISTS releases (
		id INTEGER PRIMARY KEY,
		artist_id VARCHAR(255),
		title VARCHAR(255),
		release_date TEXT,
		country TEXT,
		barcode INTEGER
	);`

	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	return db, nil
}
