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

	"github.com/sci-f/scif-go/internal/pkg/logger"
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
	Base        string              // /scif is the overall base
	Data        string              // <Base>/data is the data base
	Apps        string              // <Base>/apps is the apps base
	ShellCmd    string              // default shell
	EntryPoint  []string            // default entrypoint to an app (parsed to list)
	EntryFolder string              // default entryfolder TODO: what should this be?
	allowAppend bool                // allow appending to path
	appendPaths [3]string
	scifApps    []string
	activeApp   string              // the active app (if one is defined)
        config      map[string]string   // a loaded configuration
}

// Printing
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

	data = getenv("SCIF_BASE", data)
	apps = getenv("SCIF_BASE", apps)

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
		Data:        data,
		Apps:        apps,
		ShellCmd:    shell,
		EntryPoint:  entrylist,
		EntryFolder: entryfolder,
		allowAppend: allowAppend,
		appendPaths: scifAppendPaths,
		scifApps:    scifApps}

	// Setup includes loading a scif, if found at base
	client.Setup()

	return client
}

// provide client to user as "Scif"
var Scif ScifClient = *NewScifClient()

//    def load(self, path, app=None, quiet=False):
//        '''load a scif recipe into the object

//            Parameters
//            ==========
//            path: the complete path to the config (recipe file) to load, or 
//                  root path of filesystem (that from calling function defaults to
//                  /scif)
//            app:  if running with context of an active app, this will load the
//                  active app environment for it as well.
//        '''
//        # 1. path is a recipe
//        if os.path.isfile(path):
//            self._config = load_recipe(path)

//        # 2. path is a base
//        elif os.path.isdir(path):
//            self._config = load_filesystem(path, quiet=quiet)

//        else:
//            bot.warning('%s is not detected as a recipe or base.' %path)
//            self._config = None

//        self.update_env(app)

// Commands

// Execute will execute a command to a scientific filesystem
//func (f Scif) Execute() {
//    fmt.Println("Execute() here")
//    //ScifRecipe._exec = _exec
//}

//# Environment
//ScifRecipe.append_path = append_path
//ScifRecipe._append_path = get_append_path
//ScifRecipe._init_env = init_env
//ScifRecipe.add_env = add_env
//ScifRecipe.export_env = export_env
//ScifRecipe.get_env = get_env
//ScifRecipe.load_env = load_env
//ScifRecipe.update_env = update_env

// Preview
//ScifRecipe.preview = preview
//ScifRecipe._preview_base = preview_base
//ScifRecipe._preview_apps = preview_apps
//ScifRecipe._init_app_preview = init_app_preview
//ScifRecipe._preview_runscript = preview_runscript
//ScifRecipe._preview_labels = preview_labels
//ScifRecipe._preview_environment = preview_environment
//ScifRecipe._preview_commands = preview_commands
//ScifRecipe._preview_files = preview_files
//ScifRecipe._preview_recipe = preview_recipe
//ScifRecipe._preview_test = preview_test

// Setup
//ScifRecipe._install_base = install_base
//ScifRecipe.set_base = set_base
//ScifRecipe.set_defaults = set_defaults

// Apps
//ScifRecipe.get_appenv_lookup = get_appenv_lookup
//ScifRecipe.get_appenv = get_appenv
//ScifRecipe.app = app
//ScifRecipe.apps = apps
//ScifRecipe.activate = activate
//ScifRecipe.deactivate = deactivate
//ScifRecipe.inspect = inspect
//ScifRecipe.reset = reset

// Installation
//ScifRecipe.install = install
//ScifRecipe._init_app = init_app
//ScifRecipe._install_apps = install_apps
//ScifRecipe._install_commands = install_commands
//ScifRecipe._install_environment = install_environment
//ScifRecipe._install_files = install_files
//ScifRecipe._install_help = install_help
//ScifRecipe._install_labels = install_labels
//ScifRecipe._install_recipe = install_recipe
//ScifRecipe._install_runscript = install_runscript
//ScifRecipe._install_script = install_script
//ScifRecipe._install_test = install_test

// Execute will execute a command to a scientific filesystem
func (client ScifClient) Execute() {
	logger.Debugf("Execute() here")
	//ScifRecipe._exec = _exec
}

// Run will run a scientific application runscript
func (cli ScifClient) Run() {
	fmt.Println("Run() here")
}

// Shell will shell into a scientific filesystem
// TODO add SCIF_SHELL as envar
func (cli ScifClient) Shell() {
	fmt.Println("Shell() here")
}

// Test a scientific filesystem
func (cli ScifClient) Test() {
	fmt.Println("Test() here")
}

//# Environment
//ScifRecipe.append_path = append_path
//ScifRecipe._append_path = get_append_path
//ScifRecipe._init_env = init_env
//ScifRecipe.add_env = add_env
//ScifRecipe.export_env = export_env
//ScifRecipe.get_env = get_env
//ScifRecipe.load_env = load_env
//ScifRecipe.update_env = update_env

// Preview
//ScifRecipe.preview = preview
//ScifRecipe._preview_base = preview_base
//ScifRecipe._preview_apps = preview_apps
//ScifRecipe._init_app_preview = init_app_preview
//ScifRecipe._preview_runscript = preview_runscript
//ScifRecipe._preview_labels = preview_labels
//ScifRecipe._preview_environment = preview_environment
//ScifRecipe._preview_commands = preview_commands
//ScifRecipe._preview_files = preview_files
//ScifRecipe._preview_recipe = preview_recipe
//ScifRecipe._preview_test = preview_test

// Setup
//ScifRecipe._install_base = install_base
//ScifRecipe.set_base = set_base
//ScifRecipe.set_defaults = set_defaults

// Apps
//ScifRecipe.get_appenv_lookup = get_appenv_lookup
//ScifRecipe.get_appenv = get_appenv
//ScifRecipe.app = app
//ScifRecipe.apps = apps
//ScifRecipe.activate = activate
//ScifRecipe.deactivate = deactivate
//ScifRecipe.inspect = inspect
//ScifRecipe.reset = reset

// Installation
//ScifRecipe.install = install
//ScifRecipe._init_app = init_app
//ScifRecipe._install_apps = install_apps
//ScifRecipe._install_commands = install_commands
//ScifRecipe._install_environment = install_environment
//ScifRecipe._install_files = install_files
//ScifRecipe._install_help = install_help
//ScifRecipe._install_labels = install_labels
//ScifRecipe._install_recipe = install_recipe
//ScifRecipe._install_runscript = install_runscript
//ScifRecipe._install_script = install_script
//ScifRecipe._install_test = install_test
