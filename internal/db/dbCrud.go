package db

import (
	"database/sql"

	"github.com/Jason-Kng/Side-B/pkg/types"

	_ "github.com/mattn/go-sqlite3"
)

func AddRecord(db *sql.DB, rec types.Record) error {
	query := `INSERT INTO records (artist_id, title) VALUES (?, ?)`
	_, err := db.Exec(query, rec.Artist, rec.Title)
	return err
}

func GetAllRecords(db *sql.DB) ([]types.Record, error) {
	rows, err := db.Query("SELECT id, artist_id, title FROM records")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []types.Record

	for rows.Next() {
		var r types.Record

		err := rows.Scan(&r.ID, &r.Artist, &r.Title)
		if err != nil {
			return nil, err
		}

		records = append(records, r)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return records, nil
}
