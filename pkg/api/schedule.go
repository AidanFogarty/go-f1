package api

import (
	"context"
	"fmt"
)

// Schedule retrieves the schedule for a season in a giver `year`.
func (ergast *Ergast) Schedule(ctx context.Context, year int) ([]Race, error) {

	url := fmt.Sprintf("%s/%d", ergast.BaseURL, year)
	racetable, err := ergast.doAction(ctx, url, defaultOffset, defaultLimit)
	if err != nil {
		return nil, err
	}
	return racetable.RaceTable.Races, nil
}
