package api

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Jason-Kng/Side-B/pkg/types"
)

func printStruct(data types.Item) {
	// MarshalIndent adds whitespace and newlines
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error printing struct:", err)
		return
	}
	fmt.Println(string(b))
}

func printUserStruct(data types.User) {
	// MarshalIndent adds whitespace and newlines
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error printing struct:", err)
		return
	}
	fmt.Println(string(b))
}

func printFolderStruct(data types.Folder) {
	// MarshalIndent adds whitespace and newlines
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error printing struct:", err)
		return
	}
	fmt.Println(string(b))
}

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

		printStruct(itemData)
		// printFolderStruct(folderData)
		// printUserStruct(userData)
		// fmt.Printf(scanner.Text())
	}
}

func ReturnRecord(resp *http.Response) (*types.Record, error) {
	scanner := bufio.NewScanner(resp.Body)
	var itemData types.Item
	for scanner.Scan() {
		if err := json.Unmarshal(scanner.Bytes(), &itemData); err != nil {
			log.Printf("Failed to parse JSON: %v", err)
		}
	}

	var barcodeNum string

	for _, item := range itemData.Identifiers {
		if item.Type == "Barcode (Scanned)" {
			barcodeNum = item.Value
			break
		} else if item.Type == "Barcode" {
			barcodeNum = item.Value
			break
		}
	}

	record := types.Record{
		ID:          itemData.ID,
		ArtistID:    itemData.Artists[0].ID,
		Title:       itemData.Title,
		ReleaseDate: itemData.Released,
		Country:     itemData.Country,
		Barcode:     barcodeNum,
	}

	return &record, nil
}
