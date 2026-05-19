package ui

import (
	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/table"
	tea "charm.land/bubbletea/v2"
)

type Model struct {
	Table table.Model
	Help  help.Model
	Keys  CustomKeyMap
}

func (m Model) Init() tea.Cmd {
	return nil
}
