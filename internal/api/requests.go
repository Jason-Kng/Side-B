package api

import (
	"fmt"
	"net/http"
	"time"
)

func requestRelease(releaseID int, usrToken string) (*http.Response, error) {
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

// func printStruct(data types.Item) {
// 	// MarshalIndent adds whitespace and newlines
// 	b, err := json.MarshalIndent(data, "", "  ")
// 	if err != nil {
// 		fmt.Println("Error printing struct:", err)
// 		return
// 	}
// 	fmt.Println(string(b))
// }
//
// func printUserStruct(data types.User) {
// 	// MarshalIndent adds whitespace and newlines
// 	b, err := json.MarshalIndent(data, "", "  ")
// 	if err != nil {
// 		fmt.Println("Error printing struct:", err)
// 		return
// 	}
// 	fmt.Println(string(b))
// }
//
// func printFolderStruct(data types.Folder) {
// 	// MarshalIndent adds whitespace and newlines
// 	b, err := json.MarshalIndent(data, "", "  ")
// 	if err != nil {
// 		fmt.Println("Error printing struct:", err)
// 		return
// 	}
// 	fmt.Println(string(b))
// }
// func main() {
// 	discogsToken := "mMlciUEsBDqynVcWsBzdrftPmgQvzwlkDyxokrsM"
// 	req, err := http.NewRequest(http.MethodGet, "https://api.discogs.com/releases/37107642", nil)
// 	// req, err := http.NewRequest(http.MethodGet, "https://api.discogs.com/oauth/identity", nil)
// 	// req, err := http.NewRequest(http.MethodGet, "https://api.discogs.com/users/JKing05/collection/folders/0/releases", nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	token := fmt.Sprintf("Discogs token=%v", discogsToken)
//
// 	req.Header.Add("User-Agent", "Side-BTerminalDB/0.1")
// 	req.Header.Add("Authorization", token)
//
// 	client := &http.Client{Timeout: 10 * time.Second}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	defer resp.Body.Close() // Always close the body
//
// 	scanner := bufio.NewScanner(resp.Body)
// 	for scanner.Scan() {
// 		var itemData types.Item
// 		// var folderData types.Folder
// 		// var userData types.User
// 		if err := json.Unmarshal(scanner.Bytes(), &itemData); err != nil {
// 			log.Printf("Failed to parse JSON: %v", err)
// 		}
//
// 		// for _, release := range data.Releases {
// 		// fmt.Printf(("Release: %v\n"), release.BasicInformation.Title)
// 		// }
// 		printStruct(itemData)
// 		// printFolderStruct(folderData)
// 		// printUserStruct(userData)
// 		// fmt.Printf(scanner.Text())
// 	}
// }
