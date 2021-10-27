package api

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"
)

// Schedule retrieves the schedule for a season in a giver `year`.
func (ergast *Ergast) Schedule(ctx context.Context, year int) ([]Race, error) {

	url := fmt.Sprintf("%s/%d", ergast.BaseURL, year)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	data, err := ergast.doAction(req)
	if err != nil {
		return nil, err
	}

	racetable := new(MRData)
	if err := xml.Unmarshal(data, racetable); err != nil {
		return nil, err
	}

	return racetable.RaceTable.Races, nil
}
