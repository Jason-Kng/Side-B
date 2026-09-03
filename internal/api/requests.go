package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Jason-Kng/Side-B/pkg/logger"
	"github.com/Jason-Kng/Side-B/pkg/types"
)

func formatBarcode(barcode string) string {
	fmtBarcode := strings.ReplaceAll(barcode, " ", "")
	return fmtBarcode
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

func ReturnRecord(resp *http.Response) (*types.ReleaseData, error) {
	var itemData types.Item
	err := json.NewDecoder(resp.Body).Decode(&itemData)
	if err != nil {
		logger.Error("Decode Error %v", err)
		return nil, fmt.Errorf("decode error: %w", err)
	}

	var barcodeNum string

	for _, item := range itemData.Identifiers {
		if item.Type == "Barcode" && item.Description == "Scanned" {
			barcodeNum = item.Value
			break
		} else if item.Type == "Barcode" {
			barcodeNum = formatBarcode(item.Value)
			break
		}
	}

	var ReleaseData types.ReleaseData

	release := types.Release{
		ID:          itemData.ID,
		MasterID:    itemData.MasterID,
		Title:       itemData.Title,
		ArtistsSort: itemData.ArtistsSort,
		Year:        itemData.Year,
		ReleaseDate: itemData.Released,
		Country:     itemData.Country,
		Barcode:     barcodeNum,
		ThumbURL:    itemData.Thumb,
		// RawJSON: ,
		// FetchedAt: ,

	}

	artists := make([]types.Artist, 0, len(itemData.Artists))
	links := make([]types.ReleaseArtist, 0, len(itemData.Artists))
	seen := make(map[int]bool)

	for i, a := range itemData.Artists {
		if !seen[a.ID] {
			seen[a.ID] = true
			artists = append(artists, types.Artist{
				ID:       a.ID,
				Name:     a.Name,
				ThumbURL: a.ThumbnailURL,
			})
		}

		links = append(links, types.ReleaseArtist{
			ReleaseID: itemData.ID,
			ArtistID:  a.ID,
			Anv:       a.Anv,
			JoinStr:   a.Join,
			Position:  i,
		})
	}

	// record := types.Record{
	// ID:          itemData.ID,
	// ArtistID:    artistID,
	// Title:       itemData.Title,
	// ReleaseDate: itemData.Released,
	// Country:     itemData.Country,
	// Barcode:     barcodeNum,
	// }

	tracks := make([]types.Track, 0, len(itemData.Tracklist))
	for i, t := range itemData.Tracklist {
		tracks = append(tracks, types.Track{
			ReleaseID: itemData.ID,
			Position:  t.Position,
			Title:     t.Title,
			Duration:  t.Duration,
			TrackIDX:  i,
			Type:      t.Type,
		})
	}

	ReleaseData.Release = release
	ReleaseData.Tracks = tracks
	ReleaseData.Artists = artists
	ReleaseData.Links = links

	return &ReleaseData, nil
}
