package ui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/Jason-Kng/Side-B/internal/api"
	"github.com/Jason-Kng/Side-B/internal/db"
	"github.com/Jason-Kng/Side-B/pkg/logger"
)

func addRecord(usrToken string, releaseID int, store *db.RecordStore) tea.Msg {
	logger.Info("Adding record with ID: %d", releaseID)
	resp, err := api.RequestRelease(releaseID, usrToken)
	if err != nil {
		logger.Error("api.RequestRelease func error: %v", err)
		return ErrorMsg{Err: err}
	}

	defer resp.Body.Close()

	record, err := api.ReturnRecord(resp)
	if err != nil {
		logger.Error("api.ReturnRecord func error: %v", err)
		return ErrorMsg{Err: err}
	}

	err = store.AddRecord(record)
	if err != nil {
		logger.Error("AddRecord func error: %v", err)
		return ErrorMsg{Err: err}
	}

	logger.Info("Added Record Cleanly")

	return RecordAddedMsg{ReleaseID: releaseID}
}

func delRecord(releaseID int, store *db.RecordStore) tea.Msg {
	logger.Info("Deleting record with ID: %d", releaseID)
	delErr := store.DeleteRecord(releaseID)
	if delErr != nil {
		logger.Error("DeleteRecord func error: %v", delErr)
		return ErrorMsg{Err: delErr}
	}

	return RecordDeletedMsg{ReleaseID: releaseID}
}
