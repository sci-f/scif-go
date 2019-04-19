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

package client

import (
	"os"

	"github.com/sci-f/scif-go/internal/pkg/logger"
	"github.com/sci-f/scif-go/pkg/util"
)

// Help will print the help for an application, if it exists
func Help(args []string) (err error) {

	// Running an app means we load from the filesystem first
	cli := ScifClient{}.Load(Scif.Base)

	if len(args) == 0 {
		logger.Exitf("Please specify an application to see help for.")
	}

	name := args[0]

	// Ensure that the app exists on the filesystem
	if ok := util.Contains(name, cli.apps()); !ok {
		logger.Exitf("%v is not an installed application.", name)
	} else {

		// Get settings, look for help script
		lookup := cli.getAppenvLookup(name)

		if _, err := os.Stat(lookup["apphelp"]); os.IsNotExist(err) {
			logger.Infof("No help exists for %s", name)
		} else {
			printDefined("%apphelp", name, Scif.config[name].help)
		}
	}
	return err
}
