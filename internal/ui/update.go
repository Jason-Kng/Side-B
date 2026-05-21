package ui

import (
	"fmt"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"github.com/Jason-Kng/Side-B/pkg/logger"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	// m.Table = BuildTable(m.Store)

	// if added, ok := msg.(RecordAddedMsg); ok {
	// 	m.showPopup = false
	// 	if added.ReleaseID == 0 {
	// 		return m, nil
	// 	}
	// 	return m, func() tea.Msg {
	// 		logger.Info(fmt.Sprintf("Starting addRecord func"))
	// 		msg = addRecord(m.UsrToken, added.ReleaseID, m.Store)
	// 		return msg
	// 	}
	// }

	switch msg := msg.(type) {
	case RecordAddedMsg:
		logger.Info(fmt.Sprintf("Record added: %d\n", msg.ReleaseID))
		m.showPopup = false
		m.Table = BuildTable(m.Store)
		return m, cmd

	case tea.KeyMsg:
		switch {

		case key.Matches(msg, m.Keys.Quit):
			logger.Info(fmt.Sprintf("Quitting"))
			return m, tea.Quit

		case key.Matches(msg, m.Keys.Add):
			logger.Info(fmt.Sprintf("Trying to add a record"))
			m.showPopup = true
			m.addPopup = initAdd(m.UsrToken, m.Store)
			return m, nil
		}
	}

	if m.showPopup {
		popup, cmd := m.addPopup.Update(msg)
		m.addPopup = popup.(AddModel)
		return m, cmd
	}

	m.Table, cmd = m.Table.Update(msg)
	return m, cmd
}
