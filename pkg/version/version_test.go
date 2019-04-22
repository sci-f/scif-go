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

package version

import (
	"regexp"
	"testing"
)

// TestVersion ensures that the version string confirms to semvar
func TestVersionSemvar(t *testing.T) {

	// Matches <>.<>.<> with optional rc (release candidate)
	const SemVerRegex string = `v?([0-9]+)(\.[0-9]+)?(\.[0-9]+)?` +
		`(-([0-9A-Za-z\-]+(\.[0-9A-Za-z\-]+)*))?` +
		`(\+([0-9A-Za-z\-]+(\.[0-9A-Za-z\-]+)*))?(\.rc)?`

	// compile the regular expression
	var versionRegex *regexp.Regexp
	versionRegex = regexp.MustCompile("^" + SemVerRegex + "$")

	// Check that the version matches
	match := versionRegex.FindStringSubmatch(Version)
	if match == nil {
		t.Errorf("Version string %v does not conform to Semvar XX.XX.XX with optional .rc (release candidate).", Version)
	}
}
