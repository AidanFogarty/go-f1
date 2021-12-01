package api

import (
	"context"
	"fmt"
)

// Standings retrieves the driver standings for a given season `year`.
func (ergast *Ergast) DriverStandings(ctx context.Context, year int) ([]DriverStanding, error) {

	url := fmt.Sprintf("%s/%d/driverStandings", ergast.BaseURL, year)
	driverStandings, err := ergast.doAction(ctx, url, defaultOffset, defaultLimit)
	if err != nil {
		return nil, err
	}
	return driverStandings.StandingsTable.DriverStandings, nil
}
