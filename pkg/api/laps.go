package api

import (
	"context"
	"fmt"
)

func (ergast *Ergast) Laps(ctx context.Context, year int, round string) (interface{}, error) {
	url := fmt.Sprintf("%s/%d/%s/laps", ergast.BaseURL, year, round)

	laps := []Lap{}

	rowsObtained := 0
	totalRows := 0

	// Golang Do While Loop
	for ok := true; ok; ok = rowsObtained < totalRows {
		results, err := ergast.doAction(ctx, url, rowsObtained, defaultLimit)
		if err != nil {
			return nil, err
		}

		totalRows = results.Total
		rowsObtained += defaultLimit

		laps = append(laps, results.RaceTable.Races[0].Laps...)
	}

	return nil, nil
}

func mergeLaps() {}
