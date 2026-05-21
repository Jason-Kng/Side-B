package ui

import (
	"strconv"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/Jason-Kng/Side-B/internal/db"
)

// type OpenAddPopupMsg struct{}

type RecordAddedMsg struct {
	ReleaseID int
}

type AddModel struct {
	textInput textinput.Model

	usrToken string
	store    *db.RecordStore
}

func initAdd(usrToken string, store *db.RecordStore) AddModel {
	ti := textinput.New()
	ti.Placeholder = "Input a Record ID to add!"
	ti.Focus()
	ti.CharLimit = 50
	ti.SetWidth(25)

	return AddModel{textInput: ti, usrToken: usrToken, store: store}
}

func (m AddModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m AddModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, func() tea.Msg {
				return RecordAddedMsg{ReleaseID: 0}
			}
		case "enter":
			releaseID, _ := strconv.Atoi(m.textInput.Value())

			return m, func() tea.Msg {
				// return RecordAddedMsg{ReleaseID: releaseID}
				return addRecord(m.usrToken, releaseID, m.store)
			}
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m AddModel) View() tea.View {
	str := lipgloss.JoinVertical(
		lipgloss.Top,
		m.headerView(),
		m.textInput.View(),
		m.footerView(),
	)
	v := tea.NewView(str)
	// v.AltScreen = true
	return v
}

func (m AddModel) headerView() string { return "What record would you like to add?" }
func (m AddModel) footerView() string { return "\n(esc to quit)" }
