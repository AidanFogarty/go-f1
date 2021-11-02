package api

import (
	"context"
	"fmt"
)

func (ergast *Ergast) Laps(ctx context.Context, year int, round string) (interface{}, error) {
	url := fmt.Sprintf("%s/%d/%s/laps", ergast.BaseURL, year, round)
	results, err := ergast.doAction(ctx, url)
	if err != nil {
		return nil, err
	}

	// Only expecting a single race so need to take the first.
	if results.Offset+results.Limit < results.Total {
		fmt.Println("Need to call again")
	}

	return results.RaceTable.Races, nil
}
