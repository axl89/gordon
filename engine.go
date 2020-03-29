package gordon

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const uri = "/urls"

// Engine The Engine to query URLs to check
type Engine struct {
	url string
}

// URLs a list of URLSs
type URLs struct {
	ID         int     `json:"id"`
	URL        string  `json:"url"`
	ExpectCode int     `json:"expected_code"`
	MaxTimeOut float32 `json:"max_timeout"`
	Content    string  `json:"content"`
}

// EngineResponse represents the JSON response of the Engine
type EngineResponse struct {
	URLs []URLs `json:"urls"`
}

//GetJSON gets a JSON! YAY!
func (e Engine) GetJSON() *EngineResponse {
	client := http.Client{
		Timeout: time.Second * 2,
	}
	req, err := http.NewRequest(http.MethodGet, e.url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Gordon")
	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	body, readErr := ioutil.ReadAll(res.Body)

	if readErr != nil {
		log.Fatal(readErr)
	}

	var resp EngineResponse

	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return &resp

}
