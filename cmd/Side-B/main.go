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

	fmt.Println("Successfully connected to the database!")
}
