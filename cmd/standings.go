/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"time"

	ergast "github.com/AidanFogarty/go-ergast/pkg/api"
	"github.com/spf13/cobra"
)

func NewStandingsCmd() *cobra.Command {
	standingsCmd := &cobra.Command{
		Use:   "standings",
		Short: "Retrieve driver or constructor standings.",
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

	fmt.Println(standings)
	return nil
}
