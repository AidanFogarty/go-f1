package api

import (
	"context"
	"fmt"
)

func (ergast *Ergast) Results(ctx context.Context, year int, round string) ([]Result, error) {

	url := fmt.Sprintf("%s/%d/%s/results", ergast.BaseURL, year, round)
	results, err := ergast.doAction(ctx, url, defaultOffset, defaultLimit)
	if err != nil {
		return nil, err
	}

	// Only expecting a single race so need to take the first.
	race := results.RaceTable.Races[0]

	return race.Results, nil
}
