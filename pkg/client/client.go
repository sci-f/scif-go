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
	"fmt"
	"os"
	"path"

	"github.com/sci-f/scif-go/pkg/util"
)

// ScifClient holds scif client functions and settings
// The final client is provided as Scif. See other named files in this folder
// for functions specific to the client, and below for the init function.
//
// setup.go:    Setup() that should be called to auto load a scif
// install.go:  installation of base, apps, data folders
// defaults.go: used below to load defaults for client
type ScifClient struct {
	Base     string // /scif is the overall base
	Data     string // <Base>/data is the data base
	Apps     string // <Base>/apps is the apps base
	ShellCmd string // default shell

	EntryPoint         []string // active entrypoint to an app (parsed to list)
	EntryFolder        string   // active entryfolder
	defaultEntryPoint  []string // default entrypoint to an app (parsed to list)
	defaultEntryFolder string   // default entryfolder

	Environment map[string]string // key value pairs of current environment
	allowAppend bool              // allow appending to path
	appendPaths [3]string
	scifApps    []string
	activeApp   string                 // the active app (if one is defined)
	config      map[string]AppSettings // a loaded configuration
}

// AppSettings includes ScifClient data objects (under apps), meaning
// Env, Labels, Help, Runscript, Test, and Install.
// Each has it's own Data structure under the config["apps"]
type AppSettings struct {
	labels    []string
	environ   []string
	help      []string
	runscript []string
	test      []string
	install   []string
	files     []string
}

// String handles printing
func (client ScifClient) String() string {
	return fmt.Sprintf("[scif-client][base:%s]", Scif.Base)
}

// NewScifClient handles grabbing settings from the environment (an init)
func NewScifClient() *ScifClient {

	base := getenv("SCIF_BASE", getStringDefault("BASE"))
	scifApps := getenvNamespace("SCIF_APP")

	// Set the default apps and data (overridden if user sets)
	data := fmt.Sprintf(path.Join(base, "data"))
	apps := fmt.Sprintf(path.Join(base, "apps"))

	data = getenv("SCIF_DATA", data)
	apps = getenv("SCIF_APPS", apps)

	// Permissions
	allowAppend := getBoolEnv("SCIF_ALLOW_APPEND_PATHS", getBoolDefault("ALLOW_APPEND_PATHS"))
	scifAppendPaths := [3]string{"PYTHONPATH", "PATH", "LD_LIBRARY_PATH"}

	// Entry points
	shell := getenv("SCIF_SHELL", getStringDefault("SHELL"))
	entrypoint := getenv("SCIF_ENTRYPOINT", getStringDefault("ENTRYPOINT"))
	entryfolder := getenv("SCIF_ENTRYFOLDER", getStringDefault("ENTRYFOLDER"))
	entrylist := util.ParseEntrypoint(entrypoint)

	// Update Environment
	os.Setenv("SCIF_DATA", data)
	os.Setenv("SCIF_APPS", apps)
	os.Setenv("SCIF_BASE", base)

	// Instantiate the client
	client := &ScifClient{Base: base,
		Data:               data,
		Apps:               apps,
		ShellCmd:           shell,
		EntryPoint:         entrylist,
		EntryFolder:        entryfolder,
		defaultEntryPoint:  entrylist,
		defaultEntryFolder: entryfolder,
		allowAppend:        allowAppend,
		appendPaths:        scifAppendPaths,
		scifApps:           scifApps}

	// Additional setup could be run here
	return client
}

// Scif provide client to user as "Scif"
var Scif ScifClient = *NewScifClient()
