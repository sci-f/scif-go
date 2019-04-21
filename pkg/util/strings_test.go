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

import "testing"

// TestPrepend ensures that prepend will put an element at the front of an array
func TestPrepend(t *testing.T) {

	array := []string{"two", "three", "four"}
	updated := Prepend("one", array)

	// Ensure the first element has been prepended
	if updated[0] != "one" {
		t.Errorf("Element \"one\" was not correctly prepended, got: %v.", updated)
	}
}

// TestContains ensures that we test an array for containing a string (or not)
func TestContains(t *testing.T) {

	array := []string{"two", "three", "four"}

	if Contains("one", array) {
		t.Errorf("\"one\" is not contained in %v.", array)
	}

	if !Contains("two", array) {
		t.Errorf("\"two\" is contained in %v.", array)
	}
}
