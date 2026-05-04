package db

import (
	"database/sql"

	"github.com/Jason-Kng/Side-B/pkg/types"

	_ "github.com/mattn/go-sqlite3"
)

func AddRecord(db *sql.DB, rec types.Record) error {
	query := `INSERT INTO releases (id, artist_id, title, release_year) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(query, rec.ID, rec.ArtistID, rec.Title, rec.ReleaseYear)
	return err
}

func GetAllRecords(db *sql.DB) ([]types.Record, error) {
	rows, err := db.Query("SELECT * FROM releases")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []types.Record

	for rows.Next() {
		var r types.Record

		err := rows.Scan(&r.ID, &r.ArtistID, &r.Title, &r.ReleaseYear)
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
