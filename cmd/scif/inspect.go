// Copyright (C) 2019 Vanessa Sochat.

// This program is free software: you can redistribute it and/or modify it
// under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or (at your
// option) any later version.

// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.  See the GNU Affero General Public
// License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"github.com/sci-f/scif-go/cmd/scif/docs"
	"github.com/sci-f/scif-go/internal/pkg/logger"
	"github.com/sci-f/scif-go/pkg/client"
	"github.com/spf13/cobra"
)

var (
	InspectRunscript bool
	InspectEnv       bool
	InspectLabels    bool
	InspectAll       bool // default for inspect
	InspectInstall   bool
	InspectFiles     bool
	InspectTest      bool
)

func init() {
	InspectCmd.Flags().SetInterspersed(false)
	InspectCmd.Flags().BoolVarP(&InspectRunscript, "runscript", "r", false, "inspect the runscript for one or more scientific filesystem applications.")
	InspectCmd.Flags().BoolVarP(&InspectEnv, "environment", "e", false, "inspect the environment for one or more scientific filesystem applications")
	InspectCmd.Flags().BoolVarP(&InspectLabels, "labels", "l", false, "inspect the labels for one or more scientific filesystem applications.")
	InspectCmd.Flags().BoolVarP(&InspectAll, "all", "a", false, "inspect all attributes for one or more scientific filesystem applications.")
	InspectCmd.Flags().BoolVarP(&InspectInstall, "install", "i", false, "inspect install commands for one or more scientific filesystem applications.")
	ScifCmd.AddCommand(InspectCmd)
}

// InspectCmd: scif Inspect <appname>
var InspectCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Args:                  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {

		logger.Debugf("Inspect called with args %v", args)

		// If all are false, default to showing labels
		if !InspectRunscript && !InspectEnv && !InspectLabels && !InspectAll && !InspectInstall && !InspectFiles && !InspectTest {
			InspectLabels = true
		}

		// appname is optional, so likely args could be empty
		err := client.Inspect(InspectRunscript, InspectEnv, InspectLabels, InspectInstall, InspectFiles, InspectTest, InspectAll, args)
		if err != nil {
			logger.Exitf("%v", err)
		}
	},

	Use:     docs.InspectUse,
	Short:   docs.InspectShort,
	Long:    docs.InspectLong,
	Example: docs.InspectExample,
}
