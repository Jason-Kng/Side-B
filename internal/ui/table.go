package ui

import (
	"log"
	"strconv"

	"charm.land/bubbles/v2/table"
	"charm.land/lipgloss/v2"
	"github.com/Jason-Kng/Side-B/internal/db"
)

var columns = []table.Column{
	{Title: "Record ID", Width: 20},
	{Title: "Artist ID", Width: 20},
	{Title: "Album Title", Width: 25},
	{Title: "Release Date", Width: 20},
	{Title: "Country", Width: 20},
	{Title: "Barcode ID", Width: 25},
}

// Get all records

func BuildTable(store *db.RecordStore) table.Model {
	records, err := store.GetAllRecords()
	if err != nil {
		log.Printf("Error getting records: %v", err)
	}

	rows := make([]table.Row, len(records))
	for i, r := range records {
		rows[i] = table.Row{
			strconv.Itoa(r.ID),
			strconv.Itoa(r.ArtistID),
			r.Title,
			r.ReleaseDate,
			r.Country,
			r.Barcode,
		}
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(20),
		table.WithWidth(200),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)

	t.SetStyles(s)

	return t
}
