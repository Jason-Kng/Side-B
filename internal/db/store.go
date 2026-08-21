package db

import (
	"database/sql"
	_ "embed"
	"fmt"
)

var schemaSQL string

type RecordStore struct {
	db *sql.DB
}

func NewRecordStore(db *sql.DB) *RecordStore {
	return &RecordStore{db: db}
}

func (s *RecordStore) InitSchema() error {
	if _, err := s.db.Exec(schemaSQL); err != nil {
		return fmt.Errorf("failed to execute schema: %w", err)
	}
	return nil
}

// func (s *RecordStore) createReleasesTable() error {
// 	query := `
// 	CREATE TABLE IF NOT EXISTS releases (
// 		id INTEGER PRIMARY KEY NOT NULL,
// 		artist_id INTEGER,
// 		title TEXT,
// 		release_date TEXT,
// 		country TEXT,
// 		barcode TEXT
// 	);`
//
// 	_, err := s.db.Exec(query)
// 	if err != nil {
// 		return fmt.Errorf("failed to create Releases table: %v", err)
// 	}
//
// 	return nil
// }
