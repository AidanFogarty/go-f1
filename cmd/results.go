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
	"fmt"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	ergast "github.com/AidanFogarty/go-ergast/pkg/api"
	"github.com/AidanFogarty/go-f1/pkg/flags"
)

// resultsCmd represents the results command
var resultsCmd = &cobra.Command{
	Use:   "results",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("results called")
	},
}

func NewResultsCmd() *cobra.Command {
	resultsCmd := &cobra.Command{
		Use:   "results",
		Short: "Receive race results",
		RunE:  doResults,
	}

	resultsCmd.Flags().Int("year", time.Now().Year(), "The year to retrieve results for")

	return resultsCmd
}

func doResults(cmd *cobra.Command, args []string) error {
	year, err := cmd.Flags().GetInt("year")
	if err != nil {
		return err
	}

	results, err := ergast.New().Results(cmd.Context(), year, "last")
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Name", "Constructor", "Points"})

	for _, result := range results {
		name := fmt.Sprintf("%s %s %s", flags.GetFlag(result.Driver.Nationality), result.Driver.GivenName, result.Driver.FamilyName)

		table.Append([]string{result.Position, name, result.Constructor.Name, result.Points})
	}
	table.Render()

	return nil
}
