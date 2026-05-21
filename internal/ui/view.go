package ui

import (
	tea "charm.land/bubbletea/v2"
)

func (m Model) View() tea.View {
	content := baseStyle.Render(m.Table.View()) + "\n " + m.Help.View(m.Keys)

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
