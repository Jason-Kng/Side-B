package main

import (
	"fmt"
	"log"

	"github.com/Jason-Kng/Side-B/internal/db"
)

func main() {
	database, err := db.InitDB("./Side-B.db")
	if err != nil {
		log.Fatalf("Could not initalize database: %v", err)
	}
	defer database.Close()

	// newRecord := types.Record{
	// ID:          33652617,
	// ArtistID:    128537,
	// Title:       "On Guitar",
	// ReleaseYear: 2025,
	// }

	// err = db.AddRecord(database, newRecord)
	// if err != nil {
	// log.Printf("Error adding record: %v", err)
	// } else {
	// log.Println("Successfully added a record to the database!")
	// }

	records, err := db.GetAllRecords(database)
	if err != nil {
		log.Printf("Error getting records: %v", err)
	}

	fmt.Println("after GetAllRecords func")

	for _, value := range records {
		fmt.Println(value)
	}

	fmt.Println("Successfully connected to the database!")
}
