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
	"os/exec"

	"github.com/sci-f/scif-go/internal/pkg/logger"
	"github.com/sci-f/scif-go/pkg/util"
)

// Execute some commands to an executable. We first set the EntryPoint to be
// the executable, and the additional arguments are added by client.execute
func Execute(name string, executable string, cmd []string) (err error) {

	// Running an app means we load from the filesystem first
	cli := ScifClient{}.Load(Scif.Base)

	// Ensure that the app exists on the filesystem
	if ok := util.Contains(name, cli.apps()); !ok {
		logger.Warningf("%s is not an installed app.", name)
		return err
	}

	// Activate the app, meaning we set the environment and Scif.activeApp
	cli.activate(name)

	// Full path and existence checked by client.execute
	entrypoint := []string{executable}
	Scif.EntryPoint = entrypoint

	// Add additional args to the entrypoint
	logger.Debugf("Running app %s", name)

	return cli.execute(name, cmd)
}

// setup execute is the (private) function called by run, and client.Execute to
// execute the current EntryPoint for a particular app. If extra commands
// are provided, they are added. The environment is ready to go.
func (client ScifClient) execute(name string, cmd []string) (err error) {

	// Ensure that the app exists on the filesystem
	if ok := util.Contains(name, client.apps()); !ok {
		logger.Exitf("%s does not exist.", name)
	}

	// if args are provided, add on to Scif.EntryPoint
	if len(cmd) > 0 {
		Scif.EntryPoint = append(Scif.EntryPoint, cmd...)
		logger.Debugf("Args added to EntryPoint, %v", Scif.EntryPoint)
	}

	// Add additional args to the entrypoint
	logger.Debugf("Executing command %v for app %s", Scif.EntryPoint, name)

	// If EntryFolder still not set, just enter to base
	if Scif.EntryFolder == "" {
		Scif.EntryFolder = Scif.Base
	}

	// Change directory to the EntryFolder
	if err := os.Chdir(Scif.EntryFolder); err != nil {
		logger.Exitf("%s", err)
	}

	// Find the executable (the first in the Scif.EntryPoint)
	executable, err := exec.LookPath(Scif.EntryPoint[0])
	if err != nil {
		return err
	}

	// Commands (and args) are the remaining of the EntryPoint
	commands := Scif.EntryPoint[1:]
	commands = util.ParseEntrypointList(commands)

	logger.Infof("Executing %s:%s %v", name, executable, commands)

	// Execute the command
	process := exec.Command(executable, commands...)
	process.Stdin = os.Stdin
	process.Stdout = os.Stdout
	process.Stderr = os.Stderr
	err = process.Run()
	return err

}
