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
	//"encoding/json"
	"fmt"
	"strings"

	"github.com/sci-f/scif-go/internal/pkg/logger"
	"github.com/sci-f/scif-go/pkg/util"
	// jRes, err := util.ParseErrorBody(resp.Body)
)

// Preview an app for a scientific filesystem
// preview the complete setup for a scientific filesytem. This is useful
// to print out actions for install (without doing them).

func Preview(recipe string, apps []string) {

	logger.Debugf("Previewing recipe %s", recipe)

	// Create the client, load the recipe/filesystem (all apps included)
	cli := ScifClient{}.Load(recipe)

	// install Base folders
	cli.previewBase()
	cli.previewApps(apps)
}

// Preview Helper Functions
// these functions are added to the ScifClient struct base, and have access
// to other variables via the (initialized) Scif.<varname>)
// .............................................................................

// installBase is a private function to install the base, apps, and data folder
func (client ScifClient) previewBase() {
	logger.Infof("[base] %s", Scif.Base)
	logger.Infof("[apps] %s", Scif.Apps)
	logger.Infof("[data] %s", Scif.Data)
}

// installApps installs one or more apps to the base, apps is a list of apps.
// if Apps is an empty list (provided by the user) we by default use all those
// found in the recipe.
func (client ScifClient) previewApps(apps []string) {

	// If no apps defined, get those found at base
	if len(apps) == 0 {
		apps = client.apps()
	}

	// Loop through apps to install
	for _, app := range apps {

		// Print paths for bin, lib, etc. that aren't created
		logger.Debugf("Previewing app %s", app)
		client.printAppPreview(app)

		// Exit quickly if app isn't in the config
		if ok := util.Contains(app, client.apps()); !ok {
			logger.Exitf("App %s not found in loaded config.", app)
		}

		// Get a lookup for the folders (not created)
		lookup := client.getAppenvLookup(app)

		// Handle environment, runscript, labels
		client.previewRunscript(app, lookup)
		client.previewEnvironment(app, lookup)
		client.previewHelp(app, lookup)
		client.previewFiles(app, lookup)
		client.previewCommands(app, lookup)
		client.previewTest(app, lookup)
	}
}

// previewFiles will simply print commands that would be used for copying
func (client ScifClient) previewFiles(name string, lookup map[string]string) {

	if len(lookup["appfiles"]) > 0 {

		logger.Debugf("\n+ appfiles %s", name)
		for _, files := range lookup["appfiles"] {
			fmt.Printf("%s", files)

		}
	}
}

// preview labels for a labels
func (client ScifClient) previewLabels(name string, lookup map[string]string) {

	// Exit early if no labels
	if len(lookup["applabels"]) > 0 {

		labels := make(map[string]string)
		logger.Debugf("\n+ applabels %s", name)

		var updated, key string
		var parts []string
		for _, line := range lookup["applabels"] {

			// Split the pair by the =
			updated = strings.Replace(string(line), `=`, " ", 1)
			parts = strings.Split(updated, " ")
			key = strings.Trim(parts[0], " ")

			// Only export if value defined
			if len(parts) > 1 {
				labels[key] = strings.Trim(parts[1], " ")
			}
		}

		// Print json structure (need to test this)
		fmt.Printf("%s", labels)
	}
}

// previewCommands will show commands to install the app
func (client ScifClient) previewCommands(name string, lookup map[string]string) {

	if len(Scif.config[name].install) > 0 {
		fmt.Printf("\n+ appinstall %s", name)
		client.printScript(Scif.config[name].install, lookup["appinstall"])
	}
}

// previewRecipe: shows the content of the <name>.scif written to metadata dir
func (client ScifClient) previewRecipe(name string, lookup map[string]string) {

	var lines []string

	// Get all sections (lines) for the app (only those defined)
	lines = client.exportAppLines(name)

	// Do we have any lines to print?
	if len(lines) > 0 {
		fmt.Printf("\n+ apprecipe %s", name)
		client.printScript(lines, lookup["apprecipe"])
	}

}

// printScript is a general function used by other preview scripts
// to print the lines for a script to the terminal
func (client ScifClient) printScript(lines []string, filename string) {

	// Only install the script if the section has content
	if len(lines) > 0 {
		for _, line := range lines {
			fmt.Printf("%s\n", line)
		}
		fmt.Printf("+ %s\n", filename)
	}
}

// preview a runscript (and make executable)
func (client ScifClient) previewRunscript(name string, lookup map[string]string) {

	// Do we have any lines to print?
	if len(Scif.config[name].runscript) > 0 {
		logger.Infof("\n+ apprun %s", name)
		client.printScript(Scif.config[name].runscript, lookup["apprun"])
	}
}

// previewEnvironment: preview an environment export
func (client ScifClient) previewEnvironment(name string, lookup map[string]string) {

	if len(Scif.config[name].environ) > 0 {
		logger.Infof("\n+ appenv %s", name)
		client.printScript(Scif.config[name].environ, lookup["appenv"])
	}
}

// previewHelp to show a helpfile
func (client ScifClient) previewHelp(name string, lookup map[string]string) {
	if len(Scif.config[name].help) > 0 {
		logger.Infof("\n+ apphelp %s", name)
		client.printScript(Scif.config[name].help, lookup["apphelp"])
	}
}

// previewTest: shows a test script
func (client ScifClient) previewTest(name string, lookup map[string]string) {

	if len(Scif.config[name].test) > 0 {
		logger.Infof("\n+ apptest %s", name)
		client.printScript(Scif.config[name].test, lookup["apptest"])
	}
}
