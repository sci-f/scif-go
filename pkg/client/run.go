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
)

// Run an app for a scientific filesystem. If a user chooses
// This option, we know we are loading a Filesystem first.
func Run(name string, cmd []string) (err error) {

	// Running an app means we load from the filesystem first
	cli := ScifClient{}.Load(Scif.Base)

	// Ensure that the app exists on the filesystem
	if ok := util.Contains(name, cli.apps()); !ok {
		return err
	}

	// Activate the app, meaning we set the environment and Scif.activeApp
	cli.activate(name)

	// if args are provided, add on to Scif.EntryPoint
	if len(cmd) > 0 {
		Scif.EntryPoint = append(Scif.EntryPoint, cmd...)
		logger.Debugf("Args added to EntryPoint, %v", Scif.EntryPoint)
	}

	// Add additional args to the entrypoint
	logger.Debugf("Running app %s", name)

	return cli.execute(name)

}
