package ui

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"github.com/Jason-Kng/Side-B/pkg/logger"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case ErrorMsg:
		m.popupMode = PopupNone
		m.lastError = msg.Err

		return m, nil
	case RecordAddedMsg:
		logger.Info("Record added: %d", msg.ReleaseID)
		m.popupMode = PopupNone
		m.lastError = nil
		m.Table = BuildTable(m.Store)
		return m, nil

	case RecordDeletedMsg:
		logger.Info("Record deleted: %d", msg.ReleaseID)
		m.popupMode = PopupNone
		m.lastError = nil
		m.Table = BuildTable(m.Store)
		return m, nil

	case PopupCancelledMsg:
		m.popupMode = PopupNone
		return m, nil

	case tea.KeyMsg:
		switch {

		case key.Matches(msg, m.Keys.Quit):
			logger.Info("Quitting")
			return m, tea.Quit

		case key.Matches(msg, m.Keys.Add):
			logger.Info("Trying to add a record")
			m.popupMode = PopupAdd
			m.addPopup = initAdd(m.UsrToken, m.Store)
			return m, nil
		case key.Matches(msg, m.Keys.Delete):
			logger.Info("Trying to delete a record")
			m.popupMode = PopupDel
			m.delPopup = initDel(m.Store)
			return m, nil
		}
	}

	switch m.popupMode {
	case PopupAdd:
		popup, cmd := m.addPopup.Update(msg)
		m.addPopup = popup.(AddModel)
		return m, cmd
	case PopupDel:
		popup, cmd := m.delPopup.Update(msg)
		m.delPopup = popup.(DeleteModel)
		return m, cmd
	}

	m.Table, cmd = m.Table.Update(msg)
	return m, cmd
}
