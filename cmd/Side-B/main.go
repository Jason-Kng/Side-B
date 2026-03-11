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
	// Artist: "Charley Crockett",
	// Title:  "$10 Cowboy",
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
