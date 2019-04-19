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

// Shell into a scientific filesystem. If no args are provided, shell to
// the base. Otherwise, activate and shell to an apps base folder
func Shell(args []string) (err error) {

	// Running an app means we load from the filesystem first
	cli := ScifClient{}.Load(Scif.Base)

	if len(args) > 0 {

		name := args[0]

		// Ensure that the app exists on the filesystem
		if ok := util.Contains(name, cli.apps()); !ok {
			return err
		}

		// Activate it's environment
		cli.activate(name)

	// Otherwise, reset
	} else {
		cli.deactivate()
	}

	return cli.shell()
}

// shell is the helper function to Shell, finishing up and executing the command
// to start the shell.
func (client ScifClient) shell() (err error) {


	// If EntryFolder still not set, just enter to base
	if Scif.EntryFolder == "" {
		Scif.EntryFolder = Scif.Base
	}

	// Change directory to the EntryFolder
	if err := os.Chdir(Scif.EntryFolder); err != nil {
		logger.Exitf("%s", err)
	}

	// Find the executable (the first in the Scif.EntryPoint)
	executable, err := exec.LookPath(Scif.ShellCmd)
	if err != nil {
		return err
	}


	process := exec.Command(executable, []string{}...)
	process.Stdin = os.Stdin
	process.Stdout = os.Stdout
	process.Stderr = os.Stderr
	err = process.Run()
	return err

}
