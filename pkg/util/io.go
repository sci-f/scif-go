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
	"bufio"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/sci-f/scif-go/internal/pkg/logger"
	"golang.org/x/sys/unix"
)

// ListDirFolders returns a list of directories (one level) in a folder
func ListDirFolders(path string) []string {

	files, err := ioutil.ReadDir(path)
	if err != nil {
		logger.Exitf("%v", err)
	}

	var dirs []string

	// Add the name to the list if it's a directory
	for _, f := range files {
		if f.IsDir() {
			dirs = append(dirs, f.Name())
		}
	}
	return dirs
}

// HasWriteAccess checks if the user has write access to a path
func HasWriteAccess(path string) bool {
	return unix.Access(path, unix.W_OK) == nil
}

// WriteFile writes an array of lines to files (should include newlines)
func WriteFile(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, line := range lines {

		// Write each line (newlines already included)
		writer.WriteString(line + "\n")
	}
	return writer.Flush()
}

// ReadLines returns an array of lines from a filepath.
func ReadLines(path string) []string {

	// Read the file
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		logger.Exitf("%v", err)
	}

	// Read each line with a reader into list of lines
	var line string
	var lines []string

	reader := bufio.NewReader(file)

	for {
		line, err = reader.ReadString('\n')

		// Break when we are done
		if err != nil {
			break
		}

		// Trim the line, remove newline, add to list
		line = strings.Trim(line, "\n")
		lines = append(lines, line)
	}

	// End of file is a successful read
	if err != io.EOF {
		logger.Exitf("%v", err)
	}
	return lines
}

// WriteJson marshalls a json and writes to a file path
func WriteJson(dict map[string]string, path string) error {

	// Marshal the map into a JSON string.
	data, err := json.Marshal(dict)
	if err != nil {
		logger.Exitf("%s", err)
	}

	file, _ := json.MarshalIndent(data, "", " ")
	err = ioutil.WriteFile(path, file, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Make executable, akin to chmod u+x or chmod 0755
func MakeExecutable(path string) {

	err := os.Chmod(path, 0755)
	if err != nil {
		logger.Exitf("%s", err)
	}
}
