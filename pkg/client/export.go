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
)

// PrintConfig will print the configuration
func (client ScifClient) PrintConfig() {

	for name, settings := range Scif.config {

		printDefined("%apprun", name, settings.runscript)	
		printDefined("%appinstall", name, settings.install)	
		printDefined("%appenv", name, settings.environ)	
		printDefined("%applabels", name, settings.labels)
		printDefined("%appfiles", name, settings.files)
		printDefined("%appfiles", name, settings.help)	
		printDefined("%apptest", name, settings.test)	
	}	

}

// printIfDefined will print a section if it is non empty
func printDefined(prefix string, name string, settings []string) {
	if len(settings) > 0 {
		fmt.Println(prefix, name, "\n", strings.Join(settings, "\n"))
	}
}
