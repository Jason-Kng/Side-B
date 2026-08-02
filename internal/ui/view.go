package ui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var errorStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("196")).
	Bold(true)

func (m Model) View() tea.View {
	content := baseStyle.Render(m.Table.View()) + "\n " + m.Help.View(m.Keys)

	if m.lastError != nil {
		content += "\n" + errorStyle.Render(fmt.Sprintf("Error: %v", m.lastError))
	}

	switch m.popupMode {
	case PopupAdd:
		content += "\n" + m.addPopup.headerView() + "\n" + m.addPopup.textInput.View()
	case PopupDel:
		content += "\n" + m.delPopup.headerView() + "\n" + m.delPopup.textInput.View()
	}
	v := tea.NewView(content + "\n")
	v.AltScreen = true

	return v
}
