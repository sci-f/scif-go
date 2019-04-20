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
	"encoding/json"
	"fmt"

	"github.com/sci-f/scif-go/internal/pkg/logger"
	"github.com/sci-f/scif-go/pkg/util"
)

// Inspect one or more apps for a scientific filesystem. If None defined, inspect all.
// The boolean for "all" trumps all other settings.
func Inspect(name string, runscript bool, environ bool, labels bool, install bool, files bool, test bool, all bool, printJson bool) (err error) {

	// Running an app means we load from the filesystem first
	cli := ScifClient{}.Load(Scif.Base)

	// Ensure that the app exists on the filesystem
	if ok := util.Contains(name, cli.apps()); ok {

		// inspect the app, with Json or Not
		if printJson {
			err := cli.inspect(name, runscript, environ, labels, install, files, test, all)
			if err != nil {
				return err
			}
		} else {
			cli.inspectJson(name, runscript, environ, labels, install, files, test, all)
			if err != nil {
				return err
			}
		}
	} else {
		logger.Warningf("%s is not an installed application.", name)
	}
	return err
}

// inspect is the helper function to Inspect, finishing up and executing the command
// to start the shell.
func (client ScifClient) inspect(name string, runscript bool, environ bool, labels bool, install bool, files bool, test bool, all bool) (err error) {

	settings := Scif.config[name]

	// Keep a boolean to indicate if nothing is printed
	nothingPrinted := true

	// all trumps everything else
	if all {
		client.printAppConfig(name, settings)
	} else {

		if runscript {
			printDefined("%apphelp", name, settings.help)
			printDefined("%apprun", name, settings.runscript)
			nothingPrinted = false
		}
		if install {
			printDefined("%appinstall", name, settings.install)
			nothingPrinted = false
		}
		if labels {
			printDefined("%applabels", name, settings.labels)
			nothingPrinted = false
		}
		if environ {
			printDefined("%appenv", name, settings.environ)
			nothingPrinted = false
		}
		if files {
			printDefined("%appfiles", name, settings.files)
			nothingPrinted = false
		}
		if test {
			printDefined("%apptest", name, settings.test)
			nothingPrinted = false
		}
	}

	// Tell the user if nothing was defined
	if nothingPrinted {
		logger.Warningf("No metadata defined.")
	}
	return err
}

// inspectJson more cleanly uses the Struct to print json to the screen. We
// do this by copying the AppSettings, and then removing sections that aren't
// wanted. We use the Json API specification for formatting https://jsonapi.org/
func (client ScifClient) inspectJson(name string, runscript bool, environ bool, labels bool, install bool, files bool, test bool, all bool) (err error) {

	// Put settings into a map so we can manipulate it
	settings := make(map[string][]string)
	settings["runscript"] = Scif.config[name].runscript
	settings["install"] = Scif.config[name].install
	settings["labels"] = Scif.config[name].labels
	settings["environ"] = Scif.config[name].environ
	settings["files"] = Scif.config[name].files
	settings["test"] = Scif.config[name].test

	// Edit settings (removing those not selected) based on user selection
	if !all {

		if !runscript {
			delete(settings, "runscript")
		}
		if !install {
			delete(settings, "install")
		}
		if !labels {
			delete(settings, "labels")
		}
		if !environ {
			delete(settings, "environ")
		}
		if !files {
			delete(settings, "files")
		}
		if !test {
			delete(settings, "test")
		}
	}

	// Attributes (the app settings) are stored here
	type Attributes struct {
		Data     map[string][]string `json:"attributes"`
		dataType string              `json:"type"`
	}

	// We want to mimic the json specification for web APIs
	type JsonData struct {
		Data Attributes `json:"data"`
	}

	attributes := Attributes{Data: settings, dataType: "container"}
	data := JsonData{Data: attributes}

	result, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	// Show the result to the user
	fmt.Println(string(result))
	return err
}
