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
	"os"
	"testing"
)

// Testing Functions
//..............................................................................

// TestParseEntrypoint to test parsing the entrypoint for a scif app
func TestParseEntrypoint(t *testing.T) {

	os.Setenv("OMG", "TACOS")

	var entryPoints = []struct {
		name     string
		original string
		parsed   []string
	}{
		{"basic test", "echo hello", []string{"echo", "hello"}},
		{"[e]OMG -> $OMG", "echo [e]OMG", []string{"echo", "TACOS"}},
		{"[out] -> >", "echo vanilla [out] icecream", []string{"echo", "vanilla", ">", "icecream"}},
		{"[in] -> <", "blaa [in] bloo", []string{"blaa", "<", "bloo"}},
		{"[pipe] -> |", "cat man [pipe] grep batman", []string{"cat", "man", "|", "grep", "batman"}},
		{"[append] -> |", "cat man [append] grep batman", []string{"cat", "man", "|", "grep", "batman"}},
	}

	for _, tt := range entryPoints {
		t.Run(tt.name, func(t *testing.T) {
			entrypoint := ParseEntrypoint(tt.original)
			if !Equal(entrypoint, tt.parsed) {
				t.Errorf("got %s, want %s", entrypoint, tt.parsed)
			}
		})
	}
}

// TestParseEntrypointList to test parsing an entrypoint list
func TestParseEntrypointList(t *testing.T) {

	os.Setenv("OMG", "TACOS")

	var entryPoints = []struct {
		name     string
		original []string
		parsed   []string
	}{
		{"test basic", []string{"echo", "hello"}, []string{"echo", "hello"}},
		{"[e]OMG", []string{"echo", "[e]OMG"}, []string{"echo", "TACOS"}},
		{"[out]", []string{"[out]"}, []string{">"}},
		{"[in]", []string{"[in]"}, []string{"<"}},
		{"[pipe]", []string{"[pipe]"}, []string{"|"}},
		{"[append]", []string{"[append]"}, []string{"|"}},
	}

	for _, tt := range entryPoints {
		t.Run(tt.name, func(t *testing.T) {
			entrypoint := ParseEntrypointList(tt.original)
			if !Equal(entrypoint, tt.parsed) {
				t.Errorf("got %s, want %s", entrypoint, tt.parsed)
			}
		})
	}
}
