package api

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Jason-Kng/Side-B/pkg/logger"
	"github.com/Jason-Kng/Side-B/pkg/types"
)

func RequestRelease(releaseID int, usrToken string) (*http.Response, error) {
	url := fmt.Sprintf("https://api.discogs.com/releases/%v", releaseID)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "Side-BTerminalDB/0.1")
	req.Header.Add("Authorization", usrToken)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func ReturnResponse(resp *http.Response) {
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		var itemData types.Item
		// var folderData types.Folder
		// var userData types.User

		if err := json.Unmarshal(scanner.Bytes(), &itemData); err != nil {
			log.Printf("Failed to parse JSON: %v", err)
		}

		// printStruct(itemData)
		// printFolderStruct(folderData)
		// printUserStruct(userData)
		// fmt.Printf(scanner.Text())
	}
}

func ReturnRecord(resp *http.Response) (*types.Record, error) {
	// scanner := bufio.NewScanner(resp.Body)
	var itemData types.Item
	err := json.NewDecoder(resp.Body).Decode(&itemData)
	if err != nil {
		logger.Error("Decode Error %w", err)
		return nil, fmt.Errorf("decode error: %w", err)
	}

	var barcodeNum string

	for _, item := range itemData.Identifiers {
		if item.Type == "Barcode" && item.Description == "Scanned" {
			barcodeNum = item.Value
			break
		} else if item.Type == "Barcode" && item.Description == "" {
			barcodeNum = item.Value
			break
		}
	}

	logger.Info("Artists count: %d", len(itemData.Artists))
	if len(itemData.Artists) > 0 {
		logger.Info("First artist ID: %d", itemData.Artists[0].ID)
	}

	var artistID int

	if len(itemData.Artists) > 0 {
		artistID = itemData.Artists[0].ID
	} else {
		artistID = 0
		logger.Error("No artist found for release %d", itemData.ID)
	}

	record := types.Record{
		ID:          itemData.ID,
		ArtistID:    artistID,
		Title:       itemData.Title,
		ReleaseDate: itemData.Released,
		Country:     itemData.Country,
		Barcode:     barcodeNum,
	}

	return &record, nil
}
