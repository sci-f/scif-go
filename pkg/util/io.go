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
	"io/ioutil"
	"os"

	"github.com/sci-f/scif-go/internal/pkg/logger"
	"golang.org/x/sys/unix"

)

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
		writer.WriteString(line)
		//writer.WrieByte('\n')
	}
	return writer.Flush()
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
