package ui

import (
	"fmt"
	"strconv"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"github.com/Jason-Kng/Side-B/internal/db"
)

type RecordAddedMsg struct {
	ReleaseID int
}

type ErrorMsg struct {
	Err error
}

type PopupCancelledMsg struct{}

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

func (m AddModel) Init() tea.Cmd  { return textinput.Blink }
func (m AddModel) View() tea.View { return tea.NewView("") }

func (m AddModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, func() tea.Msg {
				return PopupCancelledMsg{}
			}
		case "enter":
			if m.textInput.Value() == "" {
				return m, func() tea.Msg {
					return PopupCancelledMsg{}
				}
			}

			releaseID, err := strconv.Atoi(m.textInput.Value())
			if err != nil {
				return m, func() tea.Msg {
					return ErrorMsg{Err: fmt.Errorf("invalid ID: %w", err)}
				}
			}

			return m, func() tea.Msg {
				return addRecord(m.usrToken, releaseID, m.store)
			}
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m AddModel) headerView() string { return "What record would you like to add?" }
func (m AddModel) footerView() string { return "\n(esc to quit)" }
