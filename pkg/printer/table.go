package printer

import (
	"fmt"
	"os"

	"github.com/AidanFogarty/go-f1/pkg/api"
	"github.com/jedib0t/go-pretty/v6/table"
)

const (
	heading = `
███████  ██████  ██████  ███    ███ ██    ██ ██       █████       ██ 
██      ██    ██ ██   ██ ████  ████ ██    ██ ██      ██   ██     ███ 
█████   ██    ██ ██████  ██ ████ ██ ██    ██ ██      ███████      ██ 
██      ██    ██ ██   ██ ██  ██  ██ ██    ██ ██      ██   ██      ██ 
██       ██████  ██   ██ ██      ██  ██████  ███████ ██   ██      ██                                                   
`
)

func ScheduleTable(races []api.Race) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.AppendHeader(table.Row{"Round", "Name", "Location", "Circuit", "Date"})

	for _, race := range races {
		t.AppendRow(table.Row{
			race.Round, race.RaceName, race.Circuit.Location.Country, race.Circuit.CircuitName, race.Date,
		})
	}
	t.Render()
}

func StandingsTable(standings []api.DriverStanding) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.AppendHeader(table.Row{"Position", "Driver", "Points", "No. Wins", "Team"})

	for _, standing := range standings {
		t.AppendRow(table.Row{
			standing.Position, standing.Driver.GivenName + " " + standing.Driver.FamilyName, standing.Points, standing.Wins, standing.Constructor.Name,
		})
	}
	t.Render()
}

func ResultsTable(results []api.Result) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.AppendHeader(table.Row{"Position", "Driver", "Nationality", "Car", "PTS"})

	for _, result := range results {
		t.AppendRow(table.Row{
			result.Position, result.Driver.GivenName + " " + result.Driver.FamilyName, result.Driver.Nationality, result.Constructor.Name, result.Points,
		})
	}
	t.Render()
}

func Heading() {
	fmt.Println(heading)
}
