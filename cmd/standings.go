/*
Copyright © 2021 Aidan Fogarty

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
	"time"

	"github.com/AidanFogarty/go-f1/pkg/api"
	"github.com/AidanFogarty/go-f1/pkg/printer"
	"github.com/spf13/cobra"
)

// standingsCmd represents the standings command.
var standingsCmd = &cobra.Command{
	Use:   "standings",
	Short: "Get the driver standings for a given year.",
	RunE: func(cmd *cobra.Command, args []string) error {
		ergast := api.New()

		standings, err := ergast.DriverStandings(cmd.Context(), year)
		if err != nil {
			return err
		}

		printer.Heading()
		printer.StandingsTable(standings)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(standingsCmd)

	standingsCmd.Flags().IntVar(&year, "year", time.Now().Year(), "The race standings year.")
}
