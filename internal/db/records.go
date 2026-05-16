package db

import "github.com/Jason-Kng/Side-B/pkg/types"

func (s *RecordStore) AddRecord(rec types.Record) error {
	query := `INSERT INTO releases (id, artist_id, title, release_date, country, barcode) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := s.db.Exec(query, rec.ID, rec.ArtistID, rec.Title, rec.ReleaseDate, rec.Country, rec.Barcode)
	return err
}

func (s *RecordStore) GetAllRecords() ([]types.Record, error) {
	rows, err := s.db.Query("SELECT * FROM releases")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []types.Record

	for rows.Next() {
		var r types.Record

		err := rows.Scan(&r.ID, &r.ArtistID, &r.Title, &r.ReleaseDate, &r.Country, &r.Barcode)
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
