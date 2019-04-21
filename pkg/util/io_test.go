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

package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// Helper Functions
//..............................................................................

// Equal is a helper function to see if two slices/arrays are equal
// This function relies on the filesystem listing being in ABC order
// (and superficially writing the directories to be in this order)
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

// CreateTempDir creates a temporary directory that is cleaned up after
func CreateTempDir(prefix string) (dir string, err error) {

	// Create directory with prefix "pizza"
	dir, err = ioutil.TempDir("", prefix)
	return dir, err
}

// Testing Functions
//..............................................................................

// TestListDirFolders will create a temporary directory with subfolders, and list
func TestListDirFolders(t *testing.T) {

	// Create directory with prefix "pizza"
	dir, err := CreateTempDir("pizza")
	if err != nil {
		t.Errorf("Error creating directory: %v", err)
	}

	// This will clean up after
	defer os.RemoveAll(dir)

	folders := []string{"cheese", "olives", "onions"}

	// Create subfolders
	for _, folder := range folders {
		path := filepath.Join(dir, folder)
		err = os.Mkdir(path, 0755)
		if err != nil {
			t.Errorf("Error creating: %v:%v", path, err)
		}
	}

	// Test listing folders
	dirs := ListDirFolders(dir)
	if !Equal(dirs, folders) {
		t.Errorf("Incorrect listing, got %v, want %v", dirs, folders)
	}
}

// TestHasWriteAccess will create a directory with write access, and then
// test that we have it.
func TestHasWriteAccess(t *testing.T) {

	// Create directory with prefix "pizza"
	dir, err := CreateTempDir("writetome")
	if err != nil {
		t.Errorf("Error creating directory: %v", err)
	}

	// This will clean up after
	defer os.RemoveAll(dir)

	if !HasWriteAccess(dir) {
		t.Errorf("%v has write access.", dir)
	}
}

// TestWriteFile tests that we can write an array of lines to file
// We also test read file :)
func TestWriteFile(t *testing.T) {

	lines := []string{"if I were", "an oscarmeyer", "weiner"}
	dir, err := CreateTempDir("dreams")
	if err != nil {
		t.Errorf("Error creating directory: %v", err)
	}

	// This will clean up after
	defer os.RemoveAll(dir)

	filename := filepath.Join(dir, "hotdogs.txt")
	if err := WriteFile(lines, filename); err != nil {
		t.Errorf("cannot write to %v:%v", filename, err)
	}

	// test the ReadLines function too!
	readlines := ReadLines(filename)
	if !Equal(lines, readlines) {
		t.Errorf("Lines read from file not equal to original, got %v, want %v", readlines, lines)
	}
}

// TestWriteJson tests marshalling a json from a dictionary.
func TestWriteJson(t *testing.T) {

	dir, err := CreateTempDir("hakuna")
	if err != nil {
		t.Errorf("Error creating directory: %v", err)
	}

	filename := filepath.Join(dir, "matata.json")

	// This will clean up after
	defer os.RemoveAll(dir)

	dict := make(map[string]string)
	dict["hash"] = "browns"
	dict["easter"] = "eggs"
	dict["chicken"] = "little"

	err = WriteJson(dict, filename)
	if err != nil {
		t.Errorf("cannot write json to %v:%v", filename, err)
	}
}

// TestMakeExecutable will ensure that a a path's permissions are changed
func TestMakeExecutable(t *testing.T) {

	file, err := ioutil.TempFile("", "lizardlips.sh")
	if err != nil {
		t.Errorf("Error creating file: %v", err)
	}

	// Get the info
	info, _ := os.Stat(file.Name())
	mode := info.Mode()
	printMode := fmt.Sprintf("%04o", mode)

	// First mode
	if printMode != "0600" {
		t.Errorf("file mode should be 0600, got: %s", printMode)
	}

	MakeExecutable(file.Name())
	info, _ = os.Stat(file.Name())
	mode = info.Mode()
	printMode = fmt.Sprintf("%04o", mode)

	// Now should be executable
	if printMode != "0755" {
		t.Errorf("file mode should be 0755, got: %s", printMode)
	}

}

//// MakeExecutable is akin to chmod u+x or chmod 0755
//func MakeExecutable(path string) {

//	err := os.Chmod(path, 0755)
//	if err != nil {
//		logger.Exitf("%s", err)
//	}
//}
