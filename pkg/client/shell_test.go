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
	"path/filepath"
	"testing"
)

// TestShell tests running a scif shell after installing to a temporary base
func TestShell(t *testing.T) {

	// Create faux scif base
	dir, err := ioutil.TempDir("", "scif")
	if err != nil {
		t.Errorf("Error creating directory: %v", err)
	}

	// This will clean up after
	defer os.RemoveAll(dir)

	// Set the base, apps, data, for testing
	Scif.Base = dir
	Scif.Apps = filepath.Join(dir, "apps")
	Scif.Data = filepath.Join(dir, "data")

	// Install recipe to the temporary base
	err = Install("../../hello-world.scif", []string{}, true)
	if err != nil {
		t.Errorf("Error installing temporary SCIF")
	}

	// Load the filesystem that was installed
	cli := ScifClient{}.Load(dir)

	// Test shell without selecting an application
	err = cli.shell()
	if err != nil {
		t.Errorf("Error running scif shell.")
	}
}
