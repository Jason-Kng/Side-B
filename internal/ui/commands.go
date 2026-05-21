package ui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/Jason-Kng/Side-B/internal/api"
	"github.com/Jason-Kng/Side-B/internal/db"
	"github.com/Jason-Kng/Side-B/pkg/logger"
)

func addRecord(usrToken string, releaseID int, store *db.RecordStore) tea.Msg {
	logger.Info(fmt.Sprintf("Adding record with ID: %d", releaseID))
	if releaseID == 0 {
		// return statusMsg(0)
	}
	resp, err := api.RequestRelease(releaseID, usrToken)
	if err != nil {
		logger.Info(fmt.Sprintf("api.RequestRelease func error: %v", err))
		// return ErrorMsg{Err: err}
	}

	defer resp.Body.Close()

	record, err := api.ReturnRecord(resp)
	if err != nil {
		logger.Info(fmt.Sprintf("api.ReturnRecord func error: %v", err))
		// return ErrorMsg{Err: err}
	}

	err = store.AddRecord(record)
	if err != nil {
		logger.Info(fmt.Sprintf("AddRecord func error: %v", err))
		// return ErrorMsg{Err: err}
	}

	logger.Info(fmt.Sprintf("Added Record Cleanly"))

	return RecordAddedMsg{ReleaseID: 0}
}
