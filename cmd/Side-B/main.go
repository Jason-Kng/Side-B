package main

import (
	"fmt"
	"log"
	"os"

	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/table"
	tea "charm.land/bubbletea/v2"
	"github.com/Jason-Kng/Side-B/internal/db"
	"github.com/Jason-Kng/Side-B/internal/ui"
	"github.com/Jason-Kng/Side-B/pkg/logger"
	"github.com/joho/godotenv"
)

func main() {
	logger.Init()
	defer logger.Close()
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

	customKM := ui.CustomKeyMap{
		KeyMap: table.DefaultKeyMap(),
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
		Add: key.NewBinding(
			key.WithKeys("a"),
			key.WithHelp("a", "add record"),
		),
		Delete: key.NewBinding(
			key.WithKeys("d"),
			key.WithHelp("d", "delete record"),
		),
	}

	m := ui.Model{
		Table: ui.BuildTable(store),
		Help:  help.New(),
		Keys:  customKM,

		Store:    store,
		UsrToken: usrToken,
	}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
