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
		artist_id VARCHAR(255),
		title VARCHAR(255),
		release_date TEXT,
		country TEXT,
		barcode INTEGER
	);`

	_, err := s.db.Exec(query)
	if err != nil {
		err := fmt.Errorf("failed to create Releases table: %v", err)
		return err
	}

	return nil
}

func (s *RecordStore) InitSchema() error {
	err := s.createReleasesTable()
	if err != nil {
		return err
	}

	return nil
}
