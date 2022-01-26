/*
Copyright © 2021 Aidan Fogarty aidan.fogarty23@mail.dcu.ie

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	ergast "github.com/AidanFogarty/go-ergast/pkg/api"
	"github.com/AidanFogarty/go-f1/pkg/flags"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func NewStandingsCmd() *cobra.Command {
	standingsCmd := &cobra.Command{
		Use:   "standings",
		Short: "Retrieve driver standings.",
		RunE:  doStandings,
	}

	standingsCmd.Flags().Int("year", time.Now().Year(), "The year to retrieve standings for")

	return standingsCmd
}

func doStandings(cmd *cobra.Command, args []string) error {
	year, err := cmd.Flags().GetInt("year")
	if err != nil {
		return err
	}

	standings, err := ergast.New().DriverStandings(cmd.Context(), year)
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Name", "Constructor", "Wins", "Points"})

	for _, standing := range standings {
		name := fmt.Sprintf("%s %s %s", flags.GetFlag(standing.Driver.Nationality), standing.Driver.GivenName, standing.Driver.FamilyName)

		table.Append([]string{standing.Position, name, standing.Constructor.Name, standing.Wins, standing.Points})
	}
	table.Render()
	return nil
}
