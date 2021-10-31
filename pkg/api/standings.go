package api

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"
)

// Standings retrieves the driver standings for a given season `year`.
func (ergast *Ergast) DriverStandings(ctx context.Context, year int) ([]DriverStanding, error) {

	url := fmt.Sprintf("%s/%d/driverStandings", ergast.BaseURL, year)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	data, err := ergast.doAction(req)
	if err != nil {
		return nil, err
	}

	driverStandings := new(MRData)
	if err := xml.Unmarshal(data, driverStandings); err != nil {
		return nil, err
	}

	return driverStandings.StandingsTable.DriverStandings, nil
}
