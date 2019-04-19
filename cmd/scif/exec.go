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
	ExecuteCmd.Flags().SetInterspersed(false)
	ScifCmd.AddCommand(ExecuteCmd)
}

// ExecCmd: scif exec <appname> <commands>
var ExecuteCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Args:                  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {

		logger.Debugf("Exec called with args %v", args)

		// If no args, exit with warning "You must supply an appname to run"
		if len(args) == 0 {
			logger.Exitf("You must supply an appname to run")
		}

		// Remove the first appname from args (pop)
		appname := args[0]
		args = args[1:]

		// If we don't have any commands to execute, no go!
		if len(args) == 0 {
			logger.Exitf("You must supply a command to run")
		}

		// The executable is now the first argument
		executable := args[0]
		args = args[1:]

		// appname string, cmd []string
		err := client.Execute(appname, executable, args); if err != nil {
			logger.Exitf("%v", err)
		}
	},
	Use:     docs.ExecUse,
	Short:   docs.ExecShort,
	Long:    docs.ExecLong,
	Example: docs.ExecExample,
}
