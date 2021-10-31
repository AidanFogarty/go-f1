package api

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	baseURL = "https://ergast.com/api/f1"
)

type MRData struct {
	XMLName        xml.Name       `xml:"MRData"`
	Total          string         `xml:"total,attr"`
	RaceTable      RaceTable      `xml:"RaceTable"`
	StandingsTable StandingsTable `xml:"StandingsTable"`
}

type RaceTable struct {
	Season string `xml:"season,attr"`
	Races  []Race `xml:"Race"`
}

type StandingsTable struct {
	Season          string           `xml:"season,attr"`
	DriverStandings []DriverStanding `xml:"StandingsList>DriverStanding"`
}

type Race struct {
	Season   string `xml:"season,attr"`
	Round    string `xml:"round,attr"`
	URL      string `xml:"url,attr"`
	RaceName string `xml:"RaceName"`
}

type DriverStanding struct {
	Position    string      `xml:"position,attr"`
	Points      string      `xml:"points,attr"`
	Wins        string      `xml:"wins,attr"`
	Driver      Driver      `xml:"Driver"`
	Constructor Constructor `xml:"Constructor"`
}

type Driver struct {
	DriverId        string `xml:"driverId,attr"`
	Code            string `xml:"code,attr"`
	URL             string `xml:"url,attr"`
	PermanentNumber string `xml:"PermanentNumber"`
	GivenName       string `xml:"GivenName"`
	FamilyName      string `xml:"FamilyName"`
	DateOfBirth     string `xml:"DateOfBirth"`
	Nationality     string `xml:"Nationality"`
}

type Constructor struct {
	ConstructorId string `xml:"constructorId,attr"`
	URL           string `xml:"url,attr"`
	Name          string `xml:"Name"`
	Nationality   string `xml:"Nationality"`
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

func (ergast *Ergast) doAction(ctx context.Context, url string) (*MRData, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

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

	mrdata := new(MRData)
	if err := xml.Unmarshal(data, mrdata); err != nil {
		return nil, err
	}

	return mrdata, nil
}
