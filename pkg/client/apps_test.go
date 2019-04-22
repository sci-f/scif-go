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
	"sort"
	"testing"
)

// TestApps to test printing all folders for a scif recipe
func TestApps(t *testing.T) {

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

	// Sort by abc order to be consistent
	sort.Strings(apps)

	// These apps should be defined
	folders := []string{"hello-custom", "hello-world-echo", "hello-world-env", "hello-world-script"}
	if !Equal(folders, apps) {
		t.Errorf("Incorrect apps listing, got %v, want %v", apps, folders)
	}
}

// TestActiveStatus tests activating and deactivating an app
func TestActiveStatus(t *testing.T) {

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

	// test not active
	testNotActive(t)

	// activate
	cli.activate("hello-custom")

	var envarsActive = []struct {
		name  string
		key   string
		value string
	}{
		{"SCIF_APPNAME", "SCIF_APPNAME", "hello-custom"},
		{"SCIF_APPENV", "SCIF_APPENV", "/scif/apps/hello-custom/scif/environment.sh"},
		{"SCIF_APPLABELS", "SCIF_APPLABELS", "/scif/apps/hello-custom/scif/labels.json"},
		{"SCIF_APPDATA", "SCIF_APPDATA", "/scif/data/hello-custom"},
		{"SCIF_APPROOT", "SCIF_APPROOT", "/scif/apps/hello-custom"},
		{"SCIF_APPHELP", "SCIF_APPHELP", "/scif/apps/hello-custom/scif/runscript.help"},
		{"SCIF_APPTEST", "SCIF_APPTEST", "/scif/apps/hello-custom/scif/test.sh"},
		{"SCIF_APPLIB", "SCIF_APPLIB", "/scif/apps/hello-custom/lib"},
		{"SCIF_APPBIN", "SCIF_APPBIN", "/scif/apps/hello-custom/bin"},
		{"SCIF_APPMETA", "SCIF_APPMETA", "/scif/apps/hello-custom/scif"},
	}

	for _, tt := range envarsActive {
		t.Run(tt.name, func(t *testing.T) {
			value := os.Getenv(tt.key)
			if value != tt.value {
				t.Errorf("got %s, want %s", value, tt.value)
			}
		})
	}

	cli.deactivate()
	testNotActive(t)

}

// testNotActive will test that an app isn't exported into environment as active
func testNotActive(t *testing.T) {

	var envarsActive = []struct {
		name  string
		key   string
		value string
	}{
		{"SCIF_APPNAME", "SCIF_APPNAME", ""},
		{"SCIF_APPENV", "SCIF_APPENV", ""},
		{"SCIF_APPLABELS", "SCIF_APPLABELS", ""},
		{"SCIF_APPDATA", "SCIF_APPDATA", ""},
		{"SCIF_APPROOT", "SCIF_APPROOT", ""},
		{"SCIF_APPHELP", "SCIF_APPHELP", ""},
		{"SCIF_APPTEST", "SCIF_APPTEST", ""},
		{"SCIF_APPLIB", "SCIF_APPLIB", ""},
		{"SCIF_APPBIN", "SCIF_APPBIN", ""},
		{"SCIF_APPMETA", "SCIF_APPMETA", ""},
	}

	for _, tt := range envarsActive {
		t.Run(tt.name, func(t *testing.T) {
			value := os.Getenv(tt.key)
			if value != tt.value {
				t.Errorf("got %s, want %s", value, tt.value)
			}
		})
	}
}

// Helper Functions
//..............................................................................

// Equal tests if two strings are equal
func Equal(one, two []string) bool {
	if len(one) != len(two) {
		return false
	}
	for i, value := range one {
		if value != two[i] {
			return false
		}
	}
	return true
}
