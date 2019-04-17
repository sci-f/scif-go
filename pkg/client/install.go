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
	"strings"

	"github.com/sci-f/scif-go/internal/pkg/logger"
	"github.com/sci-f/scif-go/pkg/util"
	// jRes, err := util.ParseErrorBody(resp.Body)
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
	if writable && !util.HasWriteAccess(Scif.Base) {
		logger.Exitf("No write access to %s", Scif.Base)
	}

	// Create the client, load the recipe/filesystem (all apps included)
	cli := ScifClient{}.Load(recipe, writable)

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
		
		// Handle environment, runscript, labels
		client.installRunscript(app, lookup)
		client.installEnvironment(app, lookup)
		client.installHelp(app, lookup)
		//TODO client.installFiles(app, lookup)
		client.installCommands(app, lookup)
		client.installRecipe(app, lookup)
		client.installTest(app, lookup)

//        self._install_test(app, settings, config)
		// After we install, in case interactive, deactivate last app
		//TODO client.deactivate(app)

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


//def install_labels(self, app, settings, config):
//    '''install labels will add labels to the app labelfile

//       Parameters
//       ==========
//       app should be the name of the app, for lookup in config['apps']
//       settings: the output of _init_app(), a dictionary of environment vars
//       config: should be the config for the app obtained with self.app(app)

//    '''
//    lookup = dict()
//    if "applabels" in config:
//        labels = config['applabels']
//        bot.level
//        bot.info('+ ' + 'applabels '.ljust(5) + app)
//        for line in labels:
//            label, value = get_parts(line, default='')
//            lookup[label] = value
//        write_json(lookup, settings['applabels'])
//    return lookup

//def install_files(self, app, settings, config):
//    '''install files will add files (or directories) to a destination.
//       If none specified, they are placed in the app base

//       Parameters
//       ==========
//       app should be the name of the app, for lookup in config['apps']
//       settings: the output of _init_app(), a dictionary of environment vars
//       config: should be the config for the app obtained with self.app(app)

//    '''
//    if "appfiles" in config:
//        files = config['appfiles']
//        bot.info('+ ' + 'appfiles '.ljust(5) + app)

//        for pair in files:
//
//            # Step 1: determine source and destination
//            src, dest = get_parts(pair, default=settings['approot'])

//            # Step 2: copy source to destination
//            cmd = ['cp']

//            if os.path.isdir(src):
//                cmd.append('-R')
//            elif os.path.exists(src):
//                cmd = cmd + [src, dest]
//                result = self._run_command(cmd)
//            else:
//                bot.warning('%s does not exist, skipping.' %src)


// install labels to a labels.json
func (client ScifClient) installLabels(name string, lookup map[string]string) {

	// Exit early if no labels
	if len(lookup["applabels"]) > 0 {


		labels := make(map[string]string)
		logger.Debugf("+ applabels %s", name)

		var updated, key string
		var parts []string
		for _, line := range(lookup["applabels"]) {

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

	if len(lookup["appinstall"]) > 0 {

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

		// Issue lines to the system (not yet tested)
		_, err = exec.Command("sh","-c", lookup["appinstall"]).Output()
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
