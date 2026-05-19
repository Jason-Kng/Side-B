package ui

import (
	"charm.land/bubbles/v2/table"
	tea "charm.land/bubbletea/v2"
)

type Model struct {
	// store *db.RecordStore
	// releases []types.Record
	Table table.Model
	// cursor int
	// err    error
}

func (m Model) Init() tea.Cmd {
	return nil
}
