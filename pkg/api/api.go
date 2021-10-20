package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	baseURL = "https://ergast.com/api/f1"
)

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

func (ergast *Ergast) doAction(request *http.Request) error {
	resp, err := ergast.HTTPClient.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("An error occured, status code: %d", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(data))

	return nil
}
