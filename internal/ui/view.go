package ui

import (
	tea "charm.land/bubbletea/v2"
)

func (m Model) View() tea.View {
	v := tea.NewView(baseStyle.Render(m.Table.View()) + "\n " + m.Help.View(m.Keys) + "\n")
	v.AltScreen = true

	return v
}
