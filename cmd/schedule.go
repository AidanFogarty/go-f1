/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	ergast "github.com/AidanFogarty/go-ergast/pkg/api"
)

func NewScheduleCmd() *cobra.Command {
	scheduleCmd := &cobra.Command{
		Use:   "schedule",
		Short: "Receive season schedule",
		RunE:  doSchedule,
	}

	scheduleCmd.Flags().Int("year", time.Now().Year(), "The year to retrieve results for")

	return scheduleCmd
}

func doSchedule(cmd *cobra.Command, args []string) error {
	year, err := cmd.Flags().GetInt("year")
	if err != nil {
		return err
	}

	schedule, err := ergast.New().Schedule(cmd.Context(), year)
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Name", "Circuit", "Date"})

	for _, race := range schedule {
		table.Append([]string{race.Round, race.RaceName, race.Circuit.CircuitName, race.Date})
	}
	table.Render()

	return nil
}
