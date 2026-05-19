package ui

import (
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/table"
)

type CustomKeyMap struct {
	table.KeyMap
	Quit key.Binding
	Add  key.Binding
}

// ShortHelp includes the quit binding
func (km CustomKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{km.LineUp, km.LineDown, km.Quit, km.Add}
}

// FullHelp includes the quit binding in the full help
func (km CustomKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{km.LineUp, km.LineDown, km.GotoTop, km.GotoBottom},
		{km.PageUp, km.PageDown, km.HalfPageUp, km.HalfPageDown},
		{km.Quit, km.Add},
	}
}
