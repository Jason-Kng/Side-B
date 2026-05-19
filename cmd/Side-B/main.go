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

	dbOut, err := store.GetAllRecords()
	if err != nil {
		log.Printf("Error getting records: %v", err)
	}

	if len(dbOut) == 0 {
		fmt.Println("The Records table is empty")
	} else {
		for _, value := range dbOut {
			fmt.Print("\n")
			fmt.Println(value)
		}
	}

	// m := ui.Model{Table: ui.BuildTable(store)}
	// if _, err := tea.NewProgram(m).Run(); err != nil {
	// 	fmt.Println("Error running program:", err)
	// 	os.Exit(1)
	// }

	quit := true
	for quit {
		var usrIn int

		fmt.Println("\nWould you like to add a new record, or view current records?")
		fmt.Println("0. Quit")
		fmt.Println("1. View Records")
		fmt.Println("2. Add New Record")
		fmt.Println("3. Delete a Record")
		fmt.Print("(input a number 0 - 3) > ")

		fmt.Scan(&usrIn)

		switch usrIn {

		case 0:
			quit = false

		case 1:
			dbOut, err := store.GetAllRecords()
			if err != nil {
				log.Printf("Error getting records: %v", err)
			}

			if len(dbOut) == 0 {
				fmt.Println("The Records table is empty")
			} else {
				for _, value := range dbOut {
					fmt.Print("\n")
					fmt.Println(value)
				}
			}
		case 2:
			var releaseIn int
			fmt.Print("Enter a record Release ID: ")

			fmt.Scan(&releaseIn)

			resp, err := api.RequestRelease(releaseIn, usrToken)
			if err != nil {
				log.Fatal("Error requesting the API")
			}

			defer resp.Body.Close()

			record, err := api.ReturnRecord(resp)
			if err != nil {
				log.Fatal(err)
			}

			err = store.AddRecord(record)
			if err != nil {
				log.Fatal(err)
			}

		case 3:
			dbOut, err := store.GetAllRecords()
			if err != nil {
				log.Printf("Error getting records: %v", err)
			}

			if len(dbOut) == 0 {
				fmt.Println("The Records table is empty")
			} else {
				for _, value := range dbOut {
					fmt.Print("\n")
					fmt.Println(value)
				}
			}

			var recID int
			fmt.Println("\nEnter the Record ID that you would like to delete.")
			fmt.Print("> ")

			fmt.Scan(&recID)

			delErr := store.DeleteRecord(recID)
			if delErr != nil {
				log.Fatal(delErr)
			}

		default:
			fmt.Println("Invalid Input Try Again!")
		}

	}
}
