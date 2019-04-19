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

func init() {
	RunCmd.Flags().SetInterspersed(false)
	ScifCmd.AddCommand(RunCmd)
}

// RunCmd: scif run <appname>
var RunCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Args:                  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {

		logger.Debugf("Run called with args %v", args)

		// If no args, exit with warning "You must supply an appname to run"
		if len(args) == 0 {
			logger.Exitf("You must supply an appname to run")
		}

		// Remove the first appname from args (pop)
		appname := args[0]
		args = args[1:]

		// appname string, cmd []string
		err := client.Run(appname, args); if err != nil {
			logger.Exitf("%v", err)
		}

		logger.Exitf("We should never get here! %v", err)
	},

	Use:     docs.RunUse,
	Short:   docs.RunShort,
	Long:    docs.RunLong,
	Example: docs.RunExample,
}
