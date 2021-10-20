package api

import (
	"context"
	"fmt"
	"net/http"
)

// Schedule retrieves the schedule for a season in a giver `year`.
func (ergast *Ergast) Schedule(ctx context.Context, year int) error {

	url := fmt.Sprintf("%s/current", ergast.BaseURL)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}

	return ergast.doAction(req)
}
