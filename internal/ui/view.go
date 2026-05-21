package ui

import (
	tea "charm.land/bubbletea/v2"
)

func (m Model) View() tea.View {
	var content string
	if m.showPopup {
		// return m.addPopup.View()
		content = baseStyle.Render(m.Table.View()) + "\n " + m.Help.View(m.Keys) + "\n" + m.addPopup.headerView() + "\n" + m.addPopup.textInput.View()
	} else {
		content = baseStyle.Render(m.Table.View()) + "\n " + m.Help.View(m.Keys)
	}
	v := tea.NewView(content + "\n")
	v.AltScreen = true

	return v
}
