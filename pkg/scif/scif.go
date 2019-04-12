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

package scif

import (
        "fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/sci-f/scif-go/docs"
	"github.com/sci-f/scif-go/internal/pkg/logger"

)

// ScifRecipe is an interface to hold a scif client functions
// cli := ScifRecipe()
// cli.Run(...) 

type ScifRecipe interface {

  // Commands
  Execute()
  Run()
  Test()
  Shell()

  // Helpers
  ////ScifRecipe._run_command = run_command
  ////ScifRecipe._set_entrypoint = set_entrypoint
  ////ScifRecipe.help = help

}
 
type Scif struct {
}


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
