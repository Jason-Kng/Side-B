package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Jason-Kng/Side-B/internal/api"
	"github.com/Jason-Kng/Side-B/pkg/types"
	"github.com/blacktop/go-termimg"
	"github.com/joho/godotenv"
)

type cover struct {
	ID       int
	ArtistID int
	thumb    string
}

func main() {
	envErr := godotenv.Load("../.env")
	if envErr != nil {
		fmt.Println("Error loading .env file")
	}

	discogsToken := os.Getenv("DISCOGS_TOKEN")

	usrToken := fmt.Sprintf("Discogs token=%v", discogsToken)

	releaseID := 37107642

	resp, err := api.RequestRelease(releaseID, usrToken)
	if err != nil {
		fmt.Println("error requesting record: %w", err)
	}

	defer resp.Body.Close()

	var itemData types.Item

	err = json.NewDecoder(resp.Body).Decode(&itemData)
	if err != nil {
		fmt.Println("Decode Error %w", err)
	}

	var artistID int

	if len(itemData.Artists) > 0 {
		artistID = itemData.Artists[0].ID
	} else {
		artistID = 0
		fmt.Printf("No artist found for release %d\n", itemData.ID)
	}

	album := cover{
		ID:       itemData.ID,
		ArtistID: artistID,
		thumb:    itemData.Thumb,
	}

	fmt.Printf("ID: %v\nArtistID: %v\nthumb: %v\n", album.ID, album.ArtistID, album.thumb)

	coverReq, _ := http.NewRequest("GET", album.thumb, nil)
	coverReq.Header.Add("User-Agent", "Side-BTerminalDB/0.1")
	coverReq.Header.Add("Authorization", usrToken)

	fmt.Printf("\nAlbum Cover Request:\nHeader: %v\nBody: %v\n", coverReq.Header, coverReq.Body)

	client := &http.Client{Timeout: 60 * time.Second}
	respCover, err := client.Do(coverReq)
	if err != nil {
		fmt.Printf("Request Error: %v\n", err)
	}

	defer respCover.Body.Close()

	if respCover.StatusCode != http.StatusOK {
		fmt.Printf("Server returned bad status: %s\n", respCover.Status)
	}

	// dump, err := httputil.DumpResponse(respCover, true)
	// if err != nil {
	// log.Fatalf("Error dumping response: %v", err)
	// }

	// 3. Print the raw byte slice as a string
	// fmt.Println(string(dump))

	outputFile := "albumCover.jpeg"
	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Failed to create file: %v\n", err)
	}
	defer file.Close()

	bytesWritten, err := io.Copy(file, respCover.Body)
	if err != nil {
		fmt.Printf("Failed to save image: %v\n", err)
	}

	fmt.Printf("\nSuccess! Saved %d bytes to %s\n", bytesWritten, outputFile)

	imgFile, err := filepath.Abs(outputFile)
	if err != nil {
		fmt.Printf("Failed to find image path: %v\n", err)
	}

	fmt.Printf("imgFile: %v\n", imgFile)

	img, err := termimg.Open(imgFile)
	if err != nil {
		fmt.Printf("Failed to open img: %v\n", err)
	}

	render, err := img.Protocol(termimg.ITerm2).Width(40).Render()
	if err != nil {
		fmt.Printf("Failed to print img: %v\n", err)
	}
	// .WidthPixels(150).HeightPixels(144)

	out := bufio.NewWriter(os.Stdout)

	// Add a couple newlines before/after to clear the prompt
	out.WriteString("\n")
	out.WriteString(render)
	out.WriteString("\n")

	// This is the crucial part: Flush() forces the data out to the terminal
	out.Flush()
}
