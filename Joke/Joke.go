package joke

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
)

// Joke struct which can house translated JSON data.
type Joke struct {
	Total  int `json:"total"`
	Result []struct {
		Category []string `json:"category"`
		IconURL  string   `json:"icon_url"`
		ID       string   `json:"id"`
		URL      string   `json:"url"`
		Value    string   `json:"value"`
	} `json:"result"`
}

// Send is public due to uppercase...
// Sends term to API
func Send(term string) Joke {
	var buffer bytes.Buffer

	const rootPath string = "https://api.chucknorris.io/jokes/"
	// https://api.chucknorris.io/jokes/categories
	endpoint := "search"
	add := "?query="
	buffer.WriteString(rootPath)
	buffer.WriteString(endpoint + add + url.QueryEscape(term))

	request, err := http.NewRequest("GET", buffer.String(), nil)
	if err != nil {
		log.Fatal("GET request url failure: ", err)
		os.Exit(1)
	}

	client := &http.Client{}

	response, error := client.Do(request)
	if error != nil {
		log.Fatal("Do: ", error)
		os.Exit(1)
	}

	defer response.Body.Close()

	var record Joke

	if err := json.NewDecoder(response.Body).Decode(&record); err != nil {
		// log.Printf("verbose error with decoding: %#v", err)
		log.Fatal("Error found in decoding: ", err)
	}

	return record
}
