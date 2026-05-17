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

	discogsToken := os.Getenv("DISCOGS_TOKEN")

	usrToken := fmt.Sprintf("Discogs token=%v", discogsToken)

	// create a connection to the sql db
	conn, err := db.Conn("./Side-B.db")
	if err != nil {
		log.Fatalf("Could not initalize database: %v", err)
	}
	defer conn.Close()

	// initalize db if doesnt exist or update db schema if it exists
	store := db.NewRecordStore(conn)
	if err := store.InitSchema(); err != nil {
		log.Fatal(err)
	}

	// records, err := store.GetAllRecords()
	// if err != nil {
	// log.Printf("Error getting records: %v", err)
	// }

	// if len(records) == 0 {
	// fmt.Println("The Records table is empty")
	// } else {
	// for _, value := range records {
	// fmt.Println(value)
	// }
	// }

	var releaseIn int

	fmt.Print("Enter a record Release ID: ")

	fmt.Scan(&releaseIn)

	resp, err := api.RequestRelease(releaseIn, usrToken) // Requesting the album "On Guitar" by Masayoshi Takanaka
	if err != nil {
		log.Fatal("Error requesting the API")
	}

	defer resp.Body.Close()

	// api.ReturnResponse(resp)
	record, err := api.ReturnRecord(resp)
	if err != nil {
		log.Fatal(err)
	}

	err = store.AddRecord(record)
	if err != nil {
		log.Fatal(err)
	}

	dbOut, err := store.GetAllRecords()
	if err != nil {
		log.Printf("Error getting records: %v", err)
	}

	if len(dbOut) == 0 {
		fmt.Println("The Records table is empty")
	} else {
		for _, value := range dbOut {
			fmt.Println(value)
		}
	}
}
