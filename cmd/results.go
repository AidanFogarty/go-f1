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
	"time"

	"github.com/AidanFogarty/go-f1/pkg/api"
	"github.com/spf13/cobra"
)

var (
	round string
)

// resultsCmd represents the results command.
var resultsCmd = &cobra.Command{
	Use:   "results",
	Short: "Get the results for a given race with year and round",
	RunE: func(cmd *cobra.Command, args []string) error {
		ergast := api.New()

		_, err := ergast.Results(cmd.Context(), year, round)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(resultsCmd)

	resultsCmd.Flags().IntVar(&year, "year", time.Now().Year(), "The race schedule year.")
	resultsCmd.Flags().StringVar(&round, "round", "last", "The race round")
}
