package ui

import (
	"fmt"
	"strconv"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"github.com/Jason-Kng/Side-B/internal/db"
)

type RecordDeletedMsg struct {
	ReleaseID int
}

type DeleteModel struct {
	textInput textinput.Model

	store *db.RecordStore
}

func initDel(store *db.RecordStore) DeleteModel {
	ti := textinput.New()
	ti.Placeholder = "Input the Record ID to delete!"
	ti.Focus()
	ti.CharLimit = 50
	ti.SetWidth(25)

	return DeleteModel{textInput: ti, store: store}
}

func (m DeleteModel) Init() tea.Cmd  { return textinput.Blink }
func (m DeleteModel) View() tea.View { return tea.NewView("") }

func (m DeleteModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
				return delRecord(releaseID, m.store)
			}

		}
	}
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m DeleteModel) headerView() string { return "What record would you like to delete?" }
func (m DeleteModel) footerView() string { return "\n(esc to quit)" }
