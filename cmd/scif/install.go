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
	"github.com/sci-f/scif-go/pkg/client" // client.Scif
	"github.com/spf13/cobra"
)

func init() {
	InstallCmd.Flags().SetInterspersed(false)
	ScifCmd.AddCommand(InstallCmd)
}

// InstallCmd is the command subgroup for scif install <recipe>
var InstallCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Args:                  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {

		// If no args, exit with warning "You must supply an appname to run"
		if len(args) == 0 {
			logger.Exitf("You must supply a recipe to install")
		}

		// Remove the first argument (the recipe)
		recipe := args[0]
		args = args[1:]

		logger.Debugf("Recipe: %v\n", recipe)
		logger.Debugf("Apps: %v\n", args)

		// recipe string, apps []string, and writable (bool)
		err := client.Install(recipe, args, !readonly)
		if err != nil {
			logger.Exitf("%v", err)
		}
	},

	Use:     docs.InstallUse,
	Short:   docs.InstallShort,
	Long:    docs.InstallLong,
	Example: docs.InstallExample,
}
