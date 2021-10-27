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

func Table(races []api.Race) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.AppendHeader(table.Row{"Round", "Name", "Season"})

	for _, race := range races {
		t.AppendRow(table.Row{
			race.Round, race.RaceName, race.Season,
		})
	}

	t.SetStyle(table.StyleColoredBlackOnRedWhite)
	t.Render()
}

func Heading() {
	fmt.Println(heading)
}
