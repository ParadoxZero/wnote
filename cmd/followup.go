// Copyright (C) 2026 Sidhin S Thomas <thomas.sidhin@outlook.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// followupCmd represents the followup command
var followupCmd = &cobra.Command{
	Use:   "followup",
	Short: "Adds a followup note",
	Long: `Add a reminder or a note to check on something later.
It automatically associates the note with the current Git repository and branch.

Example:
wn followup "Check status of gdb symbol issues in Chromium"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("followup called")
	},
}

func init() {
	rootCmd.AddCommand(followupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// followupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// followupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
