package api

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	baseURL = "https://ergast.com/api/f1"
)

type MRData struct {
	XMLName   xml.Name  `xml:"MRData"`
	Total     string    `xml:"total,attr"`
	RaceTable RaceTable `xml:"RaceTable"`
}

type RaceTable struct {
	Season string `xml:"season,attr"`
	Races  []Race `xml:"Race"`
}

type Race struct {
	Season   string `xml:"season,attr"`
	Round    string `xml:"round,attr"`
	URL      string `xml:"url,attr"`
	RaceName string `xml:"RaceName"`
}

// Ergast represents the Ergast Developer API.
type Ergast struct {
	BaseURL    string
	HTTPClient *http.Client
}

func New() *Ergast {
	return &Ergast{
		BaseURL:    baseURL,
		HTTPClient: http.DefaultClient,
	}
}

func (ergast *Ergast) doAction(req *http.Request) ([]byte, error) {
	resp, err := ergast.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() //nolint

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("An error occured, status code: %d", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
