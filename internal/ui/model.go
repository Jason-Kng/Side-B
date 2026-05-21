package ui

import (
	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/table"
	tea "charm.land/bubbletea/v2"
	"github.com/Jason-Kng/Side-B/internal/db"
)

type Model struct {
	Table table.Model
	Help  help.Model
	Keys  CustomKeyMap

	Store    *db.RecordStore
	UsrToken string

	showPopup bool
	addPopup  AddModel
	// lastError error
}

func (m Model) Init() tea.Cmd {
	return nil
}
