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
//        "os"
//	"path"

	"github.com/sci-f/scif-go/internal/pkg/logger"
//        "github.com/sci-f/scif-go/pkg/util"
)

func (client ScifClient) Setup() {
	logger.Debugf("Setup() here")
	//ScifRecipe._exec = _exec
}

// TODO: on init, need to know to read in file system (or not)
//        # If recipe path not provided, try default base
//        if path is None:
//            path = SCIF_BASE

//        # 1. Determine if path is a recipe or base
//        if path is not None:

//            self.set_base(SCIF_BASE, writable=writable) # /scif
//            self.load(path, app, quiet)                 # recipe, environment

//        # 2. Neither, development client
//        else:
//            bot.info('[skeleton] session!')
//            bot.info('           load a recipe: client.load("recipe.scif")')
//            bot.info('           change default base:  client.set_base("/")')
