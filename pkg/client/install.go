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
	"path/filepath"
	"strings"

	"github.com/sci-f/scif-go/internal/pkg/logger"
	"github.com/sci-f/scif-go/pkg/util"
)

// Install an app for a scientific filesystem
// install recipes to a base. We assume this is the root of a system
// or container, and will write the /scif directory on top of it.
// If an app name is provided, install that app if it is found
// in the config. This function goes through all steps to:
//
// 1. Install base folders to base, creating a folder for each app
// 2. Install one or more apps to it, the config is already loaded

func Install(recipe string, apps []string, writable bool) (err error) {

	logger.Debugf("Installing recipe %s", recipe)

	// Ensure that recipe exists
	if _, err := os.Stat(recipe); os.IsNotExist(err) {
		logger.Exitf("Recipe %s does not exist.", recipe)
	}

	// Ensure we have writable if asking for it
	if writable && !util.HasWriteAccess(filepath.Dir(Scif.Base)) {
		logger.Exitf("No write access to %s", Scif.Base)
	}

	// Create the client, load the recipe/filesystem (all apps included)
	cli := ScifClient{}.Load(recipe)

	// install Base folders
	cli.installBase()
	cli.installApps(apps)

	return err
}

// Install Helper Functions
// these functions are added to the ScifClient struct base, and have access
// to other variables via the (initialized) Scif.<varname>)
// .............................................................................

// installBase is a private function to install the base, apps, and data folder
func (client ScifClient) installBase() {
	logger.Infof("Installing base to %s", Scif.Base)

	// Create the base, apps folder, and data folders
	folders := []string{Scif.Base, Scif.Apps, Scif.Data}

	// Exit on any kind of error
	for _, folder := range folders {
		if err := os.MkdirAll(folder, os.ModePerm); err != nil {
			logger.Exitf("%s", err)
		}
	}
}

// installApps installs one or more apps to the base, apps is a list of apps.
// if Apps is an empty list (provided by the user) we by default use all those
// found in the recipe.
func (client ScifClient) installApps(apps []string) {

	// If no apps defined, get those found at base
	if len(apps) == 0 {
		apps = client.apps()
	}

	// Init environment for all apps
	client.initEnv(apps)

	// Loop through apps to install
	for _, app := range apps {

		// Exit quickly if app isn't in the config
		if ok := util.Contains(app, client.apps()); !ok {
			logger.Exitf("App %s not found in loaded config.", app)
		}

		logger.Infof("Installing app %s", app)

		// install the individual app (create folders)
		lookup := client.installApp(app)

		// Activate the app Environment
		client.activate(app)

		// Handle environment, runscript, labels
		client.installRunscript(app, lookup)
		client.installEnvironment(app, lookup)
		client.installHelp(app, lookup)
		client.installFiles(app, lookup)
		client.installCommands(app, lookup)
		client.installRecipe(app, lookup)
		client.installTest(app, lookup)

		// After we install deactivate last app
		client.deactivate()

	}

	// Export environment for all apps
	client.exportEnv()

}

// installApp initializes environment and installs folders for an app,
// including folders for metadata, bin, and lib to it at the SCIF_BASE
// Return appsettings so we only need to generate once
func (client ScifClient) installApp(name string) map[string]string {

	// Get a lookup for the folders
	lookup := client.getAppenvLookup(name)

	// Create these paths
	keys := []string{"appmeta", "appbin", "applib", "appdata"}

	// Exit on any kind of error
	for _, key := range keys {
		if err := os.MkdirAll(lookup[key], os.ModePerm); err != nil {
			logger.Exitf("%s", err)
		}
	}
	return lookup
}

// installFiles will copy a list of files from a source to a destination.
func (client ScifClient) installFiles(name string, lookup map[string]string) {

	if len(lookup["appfiles"]) > 0 {

		var pair []string

		logger.Debugf("+ appfiles %s", name)
		for _, files := range lookup["appfiles"] {

			cmd := []string{}

			// Split files into src and dest pairs
			pair = strings.Split(string(files), " ")

			// Handle any files not existing
			fi, err := os.Stat(pair[0])
			if err != nil {
				logger.Exitf("%s", err)
			}

			// If it's a directory, add -R for recursive
			switch mode := fi.Mode(); {
			case mode.IsDir():
				cmd = append(cmd, "-R", pair[0])
			case mode.IsRegular():
				cmd = append(cmd, pair[0])
			}

			// Add the destination
			cmd = append(cmd, pair[1])

			// Copy the source to destination, exit on fail
			_, err = exec.Command("cp", cmd...).Output()
			if err != nil {
				logger.Exitf("%s", err)
			}
			cmd = nil
		}
	}
}

// install labels to a labels.json
func (client ScifClient) installLabels(name string, lookup map[string]string) {

	// Exit early if no labels
	if len(lookup["applabels"]) > 0 {

		labels := make(map[string]string)
		logger.Debugf("+ applabels %s", name)

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

		// Write to json file
		if err := util.WriteJson(labels, lookup["applabels"]); err != nil {
			logger.Exitf("%s", err)
		}
	}
}

// install commands will finally issue commands to install the app
func (client ScifClient) installCommands(name string, lookup map[string]string) {

	if len(Scif.config[name].install) > 0 {

		logger.Debugf("+ appinstall %s", name)

		// Get the present working directory
		pwd, err := os.Getwd()
		if err != nil {
			logger.Exitf("%s", err)
		}

		// Change directory to the approot
		if err := os.Chdir(lookup["approot"]); err != nil {
			logger.Exitf("%s", err)
		}

		command := strings.Join(Scif.config[name].install, "\n")

		// Issue lines to the system (not yet tested)
		_, err = exec.Command("sh", "-c", command).Output()
		if err != nil {
			logger.Exitf("%s", err)
		}

		// Change back to pwd
		if err := os.Chdir(pwd); err != nil {
			logger.Exitf("%s", err)
		}
	}
}

// install a recipe, meaning writing the <name>.scif to the app metadata folder
func (client ScifClient) installRecipe(name string, lookup map[string]string) {

	var lines []string

	// Alert the user install the app
	logger.Debugf("+ apprecipe %s", name)

	// Get all sections (lines) for the app (only those defined)
	lines = client.exportAppLines(name)

	// The lookup contains the recipe file
	if err := util.WriteFile(lines, lookup["apprecipe"]); err != nil {
		logger.Exitf("%s", err)
	}

}

// installScript is a general function used by installRunscript, installHelp, and
// installEnvironment to write a script to a file from a config setting section
// Returns true or false if the script was written
func (client ScifClient) installScript(lines []string, filename string) bool {

	// Only install the script if the section has content
	if len(lines) > 0 {

		// Write the lines to file, if they have length
		if err := util.WriteFile(lines, filename); err != nil {
			logger.Exitf("%s", err)
		}
		return true
	}
	return false
}

// install a runscript (and make executable)
func (client ScifClient) installRunscript(name string, lookup map[string]string) {

	// Install, and then make executable (only if file exists)
	if client.installScript(Scif.config[name].runscript, lookup["apprun"]) {
		logger.Debugf("+ apprun %s", name)
		util.MakeExecutable(lookup["apprun"])
	}
}

// install an environment
func (client ScifClient) installEnvironment(name string, lookup map[string]string) {
	if client.installScript(Scif.config[name].environ, lookup["appenv"]) {
		logger.Debugf("+ appenv %s", name)
	}
}

// install a helpfile
func (client ScifClient) installHelp(name string, lookup map[string]string) {
	if client.installScript(Scif.config[name].help, lookup["apphelp"]) {
		logger.Debugf("+ apphelp %s", name)
	}
}

// install a test script
func (client ScifClient) installTest(name string, lookup map[string]string) {
	if client.installScript(Scif.config[name].test, lookup["apptest"]) {
		logger.Debugf("+ apptest %s", name)
		util.MakeExecutable(lookup["apptest"])
	}
}
