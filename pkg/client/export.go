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
	"fmt"
	"strings"

	"github.com/sci-f/scif-go/internal/pkg/logger"
)

// PrintConfig will print the configuration
func (client ScifClient) PrintConfig() {

	for name, settings := range Scif.config {

		printDefined("%apprun", name, settings.runscript)
		printDefined("%appinstall", name, settings.install)
		printDefined("%appenv", name, settings.environ)
		printDefined("%applabels", name, settings.labels)
		printDefined("%appfiles", name, settings.files)
		printDefined("%apphelp", name, settings.help)
		printDefined("%apptest", name, settings.test)
	}

}

// printIfDefined will print a section if it is non empty
func printDefined(prefix string, name string, settings []string) {
	if len(settings) > 0 {
		fmt.Println(prefix, name, "\n", strings.Join(settings, "\n"))
	}
}

// printAppPreview shows the root, lib, bin, and data for a single app
func (client ScifClient) printAppPreview(name string) {

	logger.Infof("\n\n%s", name)
	settings := client.getAppenvLookup(name)
	keys := []string{"approot", "appdata", "appbin", "applib"}

	for _, key := range keys {
		val := settings[key]
		logger.Infof("[%s] %s", key, val)
	}
}

// exportAppLines returns a list of lines for an app in a config
func (client ScifClient) exportAppLines(name string) []string {

	var lines []string
	var settings AppSettings
	var header string
	settings = Scif.config[name]

	// First add the header
	header = "%" + name + "\n"
	lines = append(lines, header)

	// Add each list of lines from the section
	lines = exportAppSection("%apprun", name, settings.runscript, lines)
	lines = exportAppSection("%appinstall", name, settings.install, lines)
	lines = exportAppSection("%appenv", name, settings.environ, lines)
	lines = exportAppSection("%applabels", name, settings.labels, lines)
	lines = exportAppSection("%appfiles", name, settings.files, lines)
	lines = exportAppSection("%apphelp", name, settings.help, lines)
	lines = exportAppSection("%apptest", name, settings.test, lines)

	return lines
}

// exportAppSection returns a list of lines for a secion
func exportAppSection(prefix string, name string, settings []string, lines []string) []string {

	var line string
	if len(settings) > 0 {
		line = prefix + " " + name + "\n"
		lines = append(lines, line)
		lines = append(lines, settings...)
	}
	return lines
}
