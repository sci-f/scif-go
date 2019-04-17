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
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/sci-f/scif-go/internal/pkg/logger"
	"github.com/sci-f/scif-go/pkg/util"
)

func (client ScifClient) Setup() {
	logger.Debugf("Running additional setup.")
}

// Loading functions. This does not coincide with doing an install (creating
// folders, etc.) but just loads a recipe or a filesystem to the config. Not
// all apps loaded in the recipe will necessarily be requested for use.
// .............................................................................

func (client ScifClient) Load(path string, writable bool) *ScifClient {

	// Initialize config and Empty environment
	Scif.config = make(map[string]AppSettings)
	Scif.Environment = make(map[string]string)

	// If the recipe is not provided (empty string) set it to be the base.
	if path == "" {
		path = Scif.Base
	}

	// Check if we have a file or a directory
	if fp, err := os.Stat(path); err == nil {

		// Case 1: It's a directory on the filesystem (scif base)
		if fp.IsDir() {

			// Load the filesystem and exit on error
			if err := client.loadFilesystem(path); err != nil {
				logger.Exitf("%s", err)
			}

		// Case 2: It's a path to a recipe
		} else {

			// Load the recipe and exit on error
			if err := client.loadRecipe(path); err != nil {
				logger.Exitf("%s", err)
			}
		}

		// Otherwise, not a recipe or directory, development mode
	} else {
		logger.Warningf("No recipe or filesystem loaded.")
	}

	client.finishLoad()
	//client.PrintConfig()

	logger.Infof("Found apps %s", client.apps())

	return &client
}

// loadRecipe is called on Load() if the path provided is a recipe file. It
// should populate the Scif.config structs
// .............................................................................

func (client ScifClient) loadRecipe(path string) error {
	logger.Infof("Calling loadRecipe, recipe %s", path)

	// Exit quickly if file doesn't exist
	if _, err := os.Stat(path); err != nil {
		return err
	}

	// Read the file
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return err
	}

	// Read each line with a reader into list of lines
	var line string
	var lines []string

	reader := bufio.NewReader(file)

	for {
		line, err = reader.ReadString('\n')

		// Break when we are done
		if err != nil {
			break
		}

		// Trim the line, remove newline, add to list
		line = strings.Trim(line, "\n")
		lines = append(lines, line)
	}

	// End of file is a successful read
	if err != io.EOF {
		return err
	}

	// We now need to populate lines into Scif.config
	section := ""
	name := ""
	var parts []string

	// Process each line
	for len(lines) > 0 {

		// Pop the first off the array
		line, lines = lines[0], lines[1:]

		// Skip comments
		if strings.HasPrefix(line, "#") {
			continue

			// A New Section
		} else if strings.HasPrefix(line, "%") {

			// Remove comments
			line = strings.Split(line, "#")[0]

			// Is there a section name?
			parts = strings.Split(line, " ")
			if len(parts) > 1 {
				name = strings.Join(parts[1:], " ")
				logger.Debugf("Found new section name %s", name)
			}

			// The section is the first part, minus the %, must be lowercase
			section = strings.Replace(parts[0], "%", "", 2)
			section = strings.ToLower(section)
			logger.Debugf("Found new section type %s", section)

			// Initialize sections for the new app (name) to Scif.config
			addSettings(name)

			// If we already have a section, we are adding to it
		} else if section != "" {

			// Add the line back to parse the section to Scif.config
			lines = util.Prepend(line, lines)
			lines = readSection(lines, section, name)
		}
	}

	// No error, woohoo!
	return nil
}

// add a new settings section, calls getSettings (but doesn't return them)
// Resulting data structure is Self.config[name]AppSettings
func addSettings(name string) {
	getSettings(name)
}

func getSettings(name string) AppSettings {

	// If the config doesn't contain apps lookup, add it
	settings, found := Scif.config[name]

	// If not found, create map subtypes
	if !found {
		settings = AppSettings{}
	}
	Scif.config[name] = settings
	return settings
}



// Read a section into Scif.config, stop when we hit the next section
func readSection(lines []string, section string, name string) []string {

	// If the config doesn't contain apps lookup, add it
	settings := getSettings(name)

	// Current members of the section will be added here
	var members []string
	var nextLine string

	for {
		// If the lines are empty, break
		if len(lines) == 0 {
			break
		}

		// Peek at the next line, don't remove from array
		nextLine = lines[0]

		// Check if the next line is a new section
		if strings.HasPrefix(nextLine, "%") {
			break

		} 	

		// Otherwise, add the nextLine to members (now remove)
                lines = lines[1:]

		// If it's not a comment
		if !strings.HasPrefix(nextLine, "#") {
			members = append(members, nextLine)
		}
	}

	// Add the list to the config
	if len(members) > 0 {
		if section != "" && name != "" {
			
			// The section determines the kind of addition we do
			switch section {
			case "appenv":
				settings.environ = members
			case "appinstall":
				settings.install = members
			case "apphelp":
				settings.help = members
			case "apprun":
				settings.runscript = members
			case "apptest":
				settings.test = members
			case "appfiles":
				settings.files = members
			case "applabels":
				settings.labels = members
			default:
				logger.Warningf("%s is not a valid section, skipping", section)
			}
		}
	}

	// Update the settings for the particular app, return smaller list lines
	Scif.config[name] = settings
	return lines
}

// loadFilesystem is called if the path provided is a Scif base (directory)
func (client ScifClient) loadFilesystem(path string) error {
	logger.Debugf("path %s", path)
	// TODO this should load the filesystem as self.config
	//             self._config = load_recipe(path)
	//def load_filesystem(base, quiet=False):
	//    '''load a filesystem based on a root path, which is usually /scif

	//        Parameters
	//        ==========
	//        base: base to load.

	//        Returns
	//        =======
	//        config: a parsed recipe configuration for SCIF
	//    '''
	//    from scif.defaults import SCIF_APPS

	//    if os.path.exists(SCIF_APPS):
	//        apps = os.listdir(SCIF_APPS)
	//        config = {'apps': {}}
	//        for app in apps:
	//            path = '%s/%s/scif/%s.scif' %(SCIF_APPS, app, app)
	//            if os.path.exists(path):
	//                recipe = load_recipe(path)
	//                config['apps'][app] = recipe['apps'][app]

	//        if len(config['apps']) > 0:
	//            if quiet is False:
	//                bot.info('Found configurations for %s scif apps' %len(config['apps']))
	//                bot.info('\n'.join(list(config['apps'].keys())))
	//            return config
	return nil
}


// finish load includes final steps to add to the runtime for an app.
// Currently, this just means adding a command to source an environment
// before running, if appenv is defined. The client should handle putting
// variables in the environment, however in some cases (if the variable
// includes an environment variable: VARIABLE1=$VARIABLE2
// It would not be properly sourced! So we add a source as the first
// line of the runscript
func (client ScifClient) finishLoad() {

	var appenv, apptest, apprun []string
	settings := AppSettings{}

	for _, app := range client.apps() {

	        // If an appenv is present for the application
		if len(Scif.config[app].environ) > 0 {

			settings = Scif.config[app]
			appenv = Scif.config[app].environ			

			// If test is defined, add source to first line
			if len(Scif.config[app].test) > 0 {
				apptest = Scif.config[app].test
				settings.test = append(apptest, appenv...)
			}

			// If runscript is defined, add source to first line
			if len(Scif.config[app].runscript) > 0 {
				apprun = Scif.config[app].runscript
				settings.runscript = append(apprun, appenv...)
			}

			Scif.config[app] = settings
		}
	}
}
