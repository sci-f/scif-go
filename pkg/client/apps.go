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

// Return a list of apps installed
func (client ScifClient) apps() []string {

	apps := make([]string, 0, len(Scif.config))
	for app := range Scif.config {
		apps = append(apps, app)
	}
	return apps
}

// activate will deactivate all apps, activate the one specified as name.
func (client ScifClient) activate(name string) {

	// Defines Scif.environment to include all vars, with name as active
	client.setActiveAppEnv(name)

	// export the changes
	client.exportEnv()
}

// deactivate will deactivate all apps
func (client ScifClient) deactivate() {

	// Reset environments for all apps (no active)
	client.initEnv(client.apps())

	// export the changes
	client.exportEnv()
}


