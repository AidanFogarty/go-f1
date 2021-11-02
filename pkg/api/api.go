package api

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	baseURL = "https://ergast.com/api/f1"
)

type MRData struct {
	XMLName        xml.Name       `xml:"MRData"`
	Total          int            `xml:"total,attr"`
	RaceTable      RaceTable      `xml:"RaceTable"`
	StandingsTable StandingsTable `xml:"StandingsTable"`
	Limit          int            `xml:"limit,attr"`
	Offset         int            `xml:"offset,attr"`
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
	Season   string   `xml:"season,attr"`
	Round    string   `xml:"round,attr"`
	URL      string   `xml:"url,attr"`
	RaceName string   `xml:"RaceName"`
	Circuit  Circuit  `xml:"Circuit"`
	Date     string   `xml:"Date"`
	Time     string   `xml:"Time"`
	Results  []Result `xml:"ResultsList>Result"`
}

type Result struct {
	Number      string      `xml:"number,attr"`
	Position    string      `xml:"position,attr"`
	Points      string      `xml:"points,attr"`
	Driver      Driver      `xml:"Driver"`
	Constructor Constructor `xml:"Constructor"`
}

type Circuit struct {
	CircuitID   string   `xml:"circuitId,attr"`
	URL         string   `xml:"url,attr"`
	CircuitName string   `xml:"CircuitName"`
	Location    Location `xml:"Location"`
}

type Location struct {
	Lat      string `xml:"lat"`
	Long     string `xml:"long"`
	Locality string `xml:"Locality"`
	Country  string `xml:"Country"`
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

func (ergast *Ergast) doAction(ctx context.Context, endpoint string) (*MRData, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	params := url.Values{}
	params.Add("limit", "105")
	req.URL.RawQuery = params.Encode()

	resp, err := ergast.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() //nolint

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("an error occured, status code: %d", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(data))
	mrdata := new(MRData)
	if err := xml.Unmarshal(data, mrdata); err != nil {
		return nil, err
	}

	return mrdata, nil
}
