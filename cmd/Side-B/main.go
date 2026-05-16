package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Jason-Kng/Side-B/internal/api"
	"github.com/Jason-Kng/Side-B/internal/db"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// discogsToken := "mMlciUEsBDqynVcWsBzdrftPmgQvzwlkDyxokrsM"
	//
	discogsToken := os.Getenv("DISCOGS_TOKEN")

	usrToken := fmt.Sprintf("Discogs token=%v", discogsToken)

	database, err := db.InitDB("./Side-B.db")
	if err != nil {
		log.Fatalf("Could not initalize database: %v", err)
	}
	defer database.Close()

	resp, err := api.RequestRelease(37107642, usrToken)
	if err != nil {
		log.Fatal("Error requesting the API")
	}

	defer resp.Body.Close()

	api.ReturnResponse(resp)

	// err = db.AddRecord(database, newRecord)
	// if err != nil {
	// log.Printf("Error adding record: %v", err)
	// } else {
	// log.Println("Successfully added a record to the database!")
	// }

	// records, err := db.GetAllRecords(database)
	// if err != nil {
	// log.Printf("Error getting records: %v", err)
	// }

	// fmt.Println("after GetAllRecords func")

	// for _, value := range records {
	// fmt.Println(value)
	// }

	// fmt.Println("Successfully connected to the database!")
}
