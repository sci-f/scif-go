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
	"io/ioutil"
	"os"
	"testing"
)

// TestShell tests running a scif shell
func TestShell(t *testing.T) {

	// Create faux scif base
	dir, err := ioutil.TempDir("", "scif")
	if err != nil {
		t.Errorf("Error creating directory: %v", err)
	}

	// This will clean up after
	defer os.RemoveAll(dir)

	// Set the base via an envar
	os.Setenv("SCIF_BASE", dir)

	// Create faux scif and get apps
	cli := ScifClient{}.Load("../../hello-world.scif")
	apps := cli.apps()

	// Test shell without selecting an application
	err = Shell([]string{})
	if err != nil {
		t.Errorf("Error running scif shell with no app selected")
	}

	// Test shell with selecting one
	err = Shell(apps)
	if err != nil {
		t.Errorf("Error running scif shell with app selection")
	}
}
