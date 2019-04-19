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
	HelpCmd.Flags().SetInterspersed(false)
	ScifCmd.AddCommand(HelpCmd)
}

// HelpCmd: scif help <appname>
var HelpCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Args:                  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {

		logger.Debugf("Help called with args %v", args)

		// appname is optional, so likely args could be empty
		err := client.Help(args)
		if err != nil {
			logger.Exitf("%v", err)
		}
	},

	Use:     docs.HelpUse,
	Short:   docs.HelpShort,
	Long:    docs.HelpLong,
	Example: docs.HelpExample,
}
