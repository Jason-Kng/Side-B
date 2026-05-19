package ui

import (
	tea "charm.land/bubbletea/v2"
)

func (m Model) View() tea.View {
	return tea.NewView(baseStyle.Render(m.Table.View()) + "\n " + m.Table.HelpView() + "\n")
}
