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

	"github.com/sci-f/scif-go/internal/pkg/logger"
	//util "github.com/sci-f/scif-go/pkg/util"
	// jRes, err := util.ParseErrorBody(resp.Body)
)

// Run an app for a scientific filesystem
func Run(appname string, cmd []string) (err error) {
	logger.Debugf("Downloading container from Shub")

	cli := ScifRecipe()
	logger.Debugf(cli)
	// cli.Run(...)

	return err
}
