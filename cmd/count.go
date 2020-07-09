/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"github.com/cemnura/scout/pkg/action"
	"github.com/spf13/cobra"

	"fmt"
)

const countDesc = `
This command takes a search PATTERN and counts the occurances in the desired directory.

Default directory is current directory.
`

// countCmd represents the count command
var countCmd = &cobra.Command{
	Use:   "count PATTERN [DIR]",
	Short: "Count occurrences of desired PATTERN in directory",
	Long:  countDesc,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ignoreCase, _ := cmd.PersistentFlags().GetBool("ignore-case")
		fmt.Print("ignoreCase : ", ignoreCase)
		result := action.Count(args[0], args[1], ignoreCase)
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(countCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	countCmd.PersistentFlags().BoolP("ignore-case", "i", false, "ignore case")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
}
