package api

import (
	"context"
	"fmt"
	"sort"
)

func (ergast *Ergast) Laps(ctx context.Context, year int, round string) ([]Lap, error) {
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

	mergedLaps := mergeLaps(laps)

	return mergedLaps, nil
}

func mergeLaps(laps []Lap) []Lap {

	seen := make(map[int]Lap)

	for _, lap := range laps {

		if val, ok := seen[lap.Number]; ok {
			merged := append(val.Timing, lap.Timing...)

			newLap := Lap{val.Number, merged}
			seen[lap.Number] = newLap
		} else {
			seen[lap.Number] = lap
		}
	}

	keys := make([]int, len(seen))
	i := 0
	for k := range seen {
		keys[i] = k
		i++
	}
	sort.Ints(keys)

	// To perform the opertion you want
	mergedLaps := []Lap{}
	for _, k := range keys {
		mergedLaps = append(mergedLaps, seen[k])
	}

	return mergedLaps
}
