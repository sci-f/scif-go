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

var longlist bool

func init() {
	AppsCmd.Flags().SetInterspersed(false)
	AppsCmd.Flags().BoolVarP(&longlist, "longlist", "l", false, "print app bases (longlist)")
	ScifCmd.AddCommand(AppsCmd)
}

// AppsCmd will list scif apps
var AppsCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Args:                  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {

		// appname is optional, so likely args could be empty
		err := client.Apps(longlist)
		if err != nil {
			logger.Exitf("%v", err)
		}
	},

	Use:     docs.AppsUse,
	Short:   docs.AppsShort,
	Long:    docs.AppsLong,
	Example: docs.AppsExample,
}
