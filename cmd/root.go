/*
Package cmd controls all the cobra cli commands for go-f1.

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
	"github.com/spf13/cobra"
)

func NewCmdRoot() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "go-f1",
		Short: "go-f1 is a cli tool to retrieve information about Formula 1 written in Go :)",
		Long: `You can use go-f1 to retrieve driver & constructor standings, race results, qualifying results,
	race lap details, pitstops and much more :)`,
	}

	rootCmd.AddCommand(NewStandingsCmd())
	rootCmd.AddCommand(NewConstructorsCmd())
	rootCmd.AddCommand(NewResultsCmd())

	return rootCmd
}
