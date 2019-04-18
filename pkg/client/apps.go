// Copyright (C) 2017-2019 Vanessa Sochat.

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
)

// Return a list of apps installed
func (client ScifClient) apps() []string {

	apps := make([]string, 0, len(Scif.config))
	for app := range Scif.config {
		apps = append(apps, app)
	}
	return apps
}

// activate will deactivate all apps, activate the one specified as name.
// We update the Scif.Environment to be relevant to the app, if one is 
// defined.
func (client ScifClient) activate(name string) {

	// deactivate any previously active apps
	client.deactivate()

	// Defines Scif.environment to include all vars, with name as active
	// This exits if the app isn't value when we call getAppenvLookup
	client.setActiveAppEnv(name)

	// Get a lookup for bin, lib, etc.
	lookup := client.getAppenvLookup(name)

	// Add bin and lib to PATH and LD_LIBRARY_PATH
	client.updatePathsFunc("PATH", lookup["appbin"])
	client.updatePathsFunc("LD_LIBRARY_PATH", lookup["applib"])

	// Set the entrypoint, if the file exists. If the user provides arguments
	// to run, these will be added by Run or Exec, etc.
	
	// If it doesn't exist, entrypoint is the default
	if _, err := os.Stat(lookup["apprun"]); os.IsNotExist(err) {
		Scif.EntryPoint = append(Scif.EntryPoint, Scif.ShellCmd)

	// Otherwise, set it to be the script
	} else {
		Scif.EntryPoint = append(Scif.EntryPoint, Scif.ShellCmd, lookup["apprun"])
	}

	logger.Debugf("EntryPoint is %v", Scif.EntryPoint)

	// Load environment variables from the app itself (environment.sh)
        client.loadAppEnv(name)

	// Set the entryfolder to the app root if it's not defined by the user
	if Scif.EntryFolder == "" {
		Scif.EntryFolder = lookup["approot"]
	}

	// Set the app to be active
	Scif.activeApp = name

	// export the changes
	client.exportEnv()

}

// deactivate will deactivate all apps
func (client ScifClient) deactivate() {

	client.activeApp = ""
	Scif.EntryFolder = Scif.defaultEntryFolder
	Scif.EntryPoint = Scif.defaultEntryPoint

	// Reset environments for all apps (no active)
	client.initEnv(client.apps())

	// export the changes
	client.exportEnv()
}


