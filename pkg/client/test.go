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
	"github.com/sci-f/scif-go/internal/pkg/logger"
	"github.com/sci-f/scif-go/pkg/util"
	"os"
)

// Test an app for a scientific filesystem. If a user chooses
// This option, we know we are loading a Filesystem first.
func Test(name string, cmd []string) (err error) {

	// Running an app means we load from the filesystem first
	cli := ScifClient{}.Load(Scif.Base)

	// Ensure that the app exists on the filesystem
	if ok := util.Contains(name, cli.apps()); !ok {
		logger.Warningf("%s is not an installed app.", name)
		return err
	}

	// Activate the app, meaning we set the environment and Scif.activeApp
	cli.activate(name)

	// Get a lookup for the folders (not created)
	lookup := cli.getAppenvLookup(name)

	// Set the entrypoint to be the test script, if it exists
	if _, err := os.Stat(lookup["apptest"]); os.IsNotExist(err) {
		logger.Warningf("No tests defined for %s", name)
		os.Exit(0)

		// Otherwise, the apptest is our entrypoint
	} else {
		Scif.EntryPoint = nil
		Scif.EntryPoint = append(Scif.EntryPoint, Scif.ShellCmd, lookup["apptest"])
	}

	// Add additional args to the entrypoint
	logger.Debugf("Testing app %s", name)

	return cli.execute(name, cmd)
}
