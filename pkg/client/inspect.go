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
	"github.com/sci-f/scif-go/pkg/util"
)

// Inspect one or more apps for a scientific filesystem. If None defined, inspect all.
// The boolean for "all" trumps all other settings.
func Inspect(runscript bool, environ bool, labels bool, install bool, files bool, test bool, all bool, apps []string) (err error) {

	// Running an app means we load from the filesystem first
	cli := ScifClient{}.Load(Scif.Base)

	// If no apps provided, default to using all
	if len(apps) == 0 {
		apps = cli.apps()
	}

	// Inspect each app
	for _, app := range apps {

		// Ensure that the app exists on the filesystem
		if ok := util.Contains(app, cli.apps()); ok {

			// inspect the app
			cli.inspect(app, runscript, environ, labels, install, files, test, all)
		}
	}
	return err
}

// shell is the helper function to Shell, finishing up and executing the command
// to start the shell.
func (client ScifClient) inspect(name string, runscript bool, environ bool, labels bool, install bool, files bool, test bool, all bool) (err error) {

	settings := Scif.config[name]

	// all trumps everything else
	if all {
		client.printAppConfig(name, settings)
	} else {

		if runscript {
			printDefined("%apphelp", name, settings.help)
			printDefined("%apprun", name, settings.runscript)
		}
		if install {
			printDefined("%appinstall", name, settings.install)
		}
		if labels {
			printDefined("%applabels", name, settings.labels)
		}
		if environ {
			printDefined("%appenv", name, settings.environ)
		}
		if files {
			printDefined("%appfiles", name, settings.files)
		}
		if test {
			printDefined("%apptest", name, settings.test)
		}
	}
	return err
}
