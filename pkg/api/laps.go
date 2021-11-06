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

	for results.Offset+results.Limit < results.Total {
		fmt.Println("Need to call again")
		results, _ = ergast.doAction(ctx, url)
		newOffset := results.Offset + results.Limit
		Offset = newOffset
	}

	return results.RaceTable.Races, nil
}
