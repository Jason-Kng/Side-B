package db

import (
	"database/sql"
	"fmt"
)

type RecordStore struct {
	db *sql.DB
}

func NewRecordStore(db *sql.DB) *RecordStore {
	return &RecordStore{db: db}
}

func (s *RecordStore) createReleasesTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS releases (
		id INTEGER PRIMARY KEY NOT NULL,
		artist_id INTEGER,
		title TEXT,
		release_date TEXT,
		country TEXT,
		barcode TEXT
	);`

	_, err := s.db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create Releases table: %v", err)
	}

	return nil
}

func (s *RecordStore) InitSchema() error {
	return s.createReleasesTable()
}
