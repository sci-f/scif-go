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
	//	"fmt"
	        "os"
	//	"path"

	"github.com/sci-f/scif-go/internal/pkg/logger"
	//        "github.com/sci-f/scif-go/pkg/util"
)

func (client ScifClient) Setup() {
	logger.Debugf("Running Setup()")
	//ScifRecipe._exec = _exec
}


func (client ScifClient) Load(path string, apps []string, writable bool) *ScifClient {
        logger.Infof("Running Load()")

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
                logger.Warningf("No recipe or filesystem loaded, development mode.")
        }

        //self.update_env(app)
        return &client
}

func (client ScifClient) loadRecipe(path string) error {
        logger.Infof("Calling loadRecipe, recipe %s", path)        
// TODO this should load the recipe as self.config
//             self._config = load_recipe(path)
        return nil
}

func (client ScifClient) loadFilesystem(path string) error {
        logger.Infof("Calling loadFilesystem, path %s", path)  
// TODO this should load the filesystem as self.config
//             self._config = load_recipe(path)
        return nil
}
