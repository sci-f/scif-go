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
	"path/filepath"
	"strings"

	"github.com/sci-f/scif-go/internal/pkg/logger"
	"github.com/sci-f/scif-go/pkg/util"
)

var envPrefix = "SCIF_"

// initEnv will initialize the complete scif environment for an app
// We parse the complete SCIF namespace from the config, and export variables
// for all apps to allow for easy interaction between them, regardless of which
// app is active. An example for a single app is provided below.
// The Scif.Environment is updated.
//
// Example: the following environment variables would be defined for an app
// 	called "google-drive" Note that for the variable, the slash is
//      replaced with an underscore

//	SCIF_APPDATA_google_drive=/scif/data/google-drive
//      SCIF_APPRUN_google_drive=/scif/apps/google-drive/scif/runscript
//      SCIF_APPHELP_google_drive=/scif/apps/google-drive/scif/runscript.help
//      SCIF_APPROOT_google_drive=/scif/apps/google-drive
//      SCIF_APPLIB_google_drive=/scif/apps/google-drive/lib
//      SCIF_APPMETA_google_drive=/scif/apps/google-drive/scif
//      SCIF_APPBIN_google_drive=/scif/apps/google-drive/bin
//      SCIF_APPENV_google_drive=/scif/apps/google-drive/scif/environment.sh
//      SCIF_APPLABELS_google_drive=/scif/apps/google-drive/scif/labels.json
//
//      These paths and files are not created at this point, but just defined.
//	A lookup for them is generated from getAppenvLookup
//
func (client ScifClient) initEnv(apps []string) {

	// if no apps provided, use those in the config
	if len(apps) == 0 {
		apps = client.apps()
	}

	// Hold all environment variables in a new map
	envars := make(map[string]string)

	// initialize base, data, and apps
	envars["SCIF_APPS"] = Scif.Apps
	envars["SCIF_BASE"] = Scif.Base
	envars["SCIF_DATA"] = Scif.Data

	// Loop through apps to export
	for _, app := range apps {

		appenv := client.getAppenvLookup(app)

		// update the values in the envars
		for k, v := range appenv {
			// SCIF_APPENV_<name>
			k = envPrefix + strings.ToUpper(k) + "_" + app
			envars[k] = v
		}
	}

	Scif.Environment = envars
}

// setActiveAppEnv sets the active app environment
func (client ScifClient) setActiveAppEnv(name string) {

	appenv := client.getAppenvLookup(name)

	// update the values in the envars
	for k, v := range appenv {
		// SCIF_APPENV_<name>
		k = envPrefix + strings.ToUpper(k)
		Scif.Environment[k] = v
	}
}

// resetEnv will reset the environment back to an empty map before updating
func (client ScifClient) resetEnv(apps []string) {
	Scif.Environment = make(map[string]string)
	client.initEnv(apps)
}

// updateEnv will update the environment, without resetting it first. It's
// equivalent to initEnv except we don't start from scratch
func (client ScifClient) updateEnv(apps []string) {

	// initialize base, data, and apps
	Scif.Environment["SCIF_APPS"] = Scif.Apps
	Scif.Environment["SCIF_BASE"] = Scif.Base
	Scif.Environment["SCIF_DATA"] = Scif.Data

	// Loop through apps to export
	for _, app := range client.apps() {

		appenv := client.getAppenvLookup(app)

		// update the values in the envars
		for k, v := range appenv {
			// SCIF_APPENV_<name>
			k = envPrefix + strings.ToUpper(k) + "_" + app
			Scif.Environment[k] = v
		}
	}
}

// appendPaths will return a string with an appended path, if allowed,
// and if the Pathname is defined in Scif.appendPaths
func (client ScifClient) appendPathsFunc(key string, value string) string {

	// If we don't allow appending, just return original value
	if !Scif.allowAppend {
		return value
	}

	// If the variable is defined on the host
	if envar, ok := os.LookupEnv(key); ok {

		// And also in the list of appendPaths
		contained := false
		for _, path := range Scif.appendPaths {
			if path == key {
				contained = true
			}
		}

		if contained {
			value = value + ":" + envar
		}
	}
	return value
}

// updatePathsFunc will call appendPathsFunc to get a new value for a path
// variable, and then set it (based on the key) into Scif.Environment
func (client ScifClient) updatePathsFunc(key string, value string) {

	value = client.appendPathsFunc(key, value)
	Scif.Environment[key] = value
}


// exportEnv will export all variables in Scif.Environment, and add the PS1
// variable by default.
func (client ScifClient) exportEnv() {

	runtime := Scif.Environment
	runtime["PS1"] = "scif> "

	// Do an update allowing extension for PATHs) and export
	for k, v := range runtime {

		// This will get any value from current env if append is allowed
		runtime[k] = client.appendPathsFunc(k, v)

		logger.Debugf("export %s=%s", k, v)
		os.Setenv(k, runtime[k])
	}
}

// loadAppEnv updates the Scif.Environment so that envars from the environment.sh
// are loaded for export when the application is activated.
func (client ScifClient) loadAppEnv(name string) {

	lookup := client.getAppenvLookup(name)

	// Determine if there is an environment.sh
	if _, err := os.Stat(lookup["appenv"]); os.IsNotExist(err) {
		return
	}

	lines := util.ReadLines(lookup["appenv"])
	var parts []string
 	
	// Parse through the lines, add them to Scif.Environment
	for _, line := range lines {
		line := strings.Trim(line, " ")
		parts = strings.Split(line, "=")

		// Skip export lines (with only one value)
		if len(parts) > 1 {
			logger.Debugf("Updating %s environment %s=%s", name, parts[0], parts[1])
			Scif.Environment[parts[0]] = parts[1]
		}
	}
}

// getAppenvLookup gets an application specific lookup for scif default
// variables. For example, an app with new "registry" would look like:
//       {'registry': {
//                      'appbin': '/scif/apps/registry/bin',
//                      'appdata': '/scif/data/registry',
//                      'appenv': '/scif/apps/registry/scif/environment.sh',
//                      'apphelp': '/scif/apps/registry/scif/runscript.help',
//                      'apptest': '/scif/apps/registry/scif/test.sh',
//                      'applabels': '/scif/apps/registry/scif/labels.json',
//                      'applib': '/scif/apps/registry/lib',
//                      'appmeta': '/scif/apps/registry/scif',
//                      'apprecipe': '/scif/apps/registry/scif/registry.scif'
//                      'approot': '/scif/apps/registry',
//                      'apprun': '/scif/apps/registry/scif/runscript'
//                    }
//       }
//       This function is intended to be shared by env above and the environment
//       generating functions in the main client, to have consistent behavior.
//       The above data structure gets parse into the (global) variables for
//       the particular app (e.g., SCIF_APPBIN_<name>
func (client ScifClient) getAppenvLookup(name string) map[string]string {

	// Exit early if app is not valid
	if ok := util.Contains(name, client.apps()); !ok {
		logger.Exitf("%s is not a valid app.", name)
	}

	envars := make(map[string]string)

	// keep the root, metadata folder, and data folder handy
	approot := filepath.Join(Scif.Apps, name) // /scif/apps/<name>
	appdata := filepath.Join(Scif.Data, name) // /scif/data/name
	appmeta := filepath.Join(approot, "scif") // /scif/apps/<name>/scif

	// Roots for app data and app files
	envars["appdata"] = appdata
	envars["approot"] = approot
	envars["appmeta"] = appmeta
	envars["appbin"] = filepath.Join(approot, "bin")
	envars["applib"] = filepath.Join(approot, "lib")
	envars["apprun"] = filepath.Join(appmeta, "runscript")
	envars["apphelp"] = filepath.Join(appmeta, "runscript.help")
	envars["appenv"] = filepath.Join(appmeta, "environment.sh")
	envars["apptest"] = filepath.Join(appmeta, "test.sh")
	envars["applabels"] = filepath.Join(appmeta, "labels.json")
	envars["apprecipe"] = filepath.Join(appmeta, name+".scif")
	envars["appname"] = name
	return envars

}
