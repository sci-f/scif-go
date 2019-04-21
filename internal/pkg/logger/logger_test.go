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

package logger

import (
	"strings"
	"testing"
)

// TestSetLevel to ensure the correct integer is returned
func TestSetLevel(t *testing.T) {
	var levelTests = []struct {
		name  string
		level int
	}{
		{"SetLevelExit", int(fatal)},
		{"SetLevelError", int(error)},
		{"SetLevelWarn", int(warn)},
		{"SetLevelLog", int(log)},
		{"SetLevelInfo", int(info)},
		{"SetLevelDebug", int(debug)},
	}

	for _, tt := range levelTests {
		t.Run(tt.name, func(t *testing.T) {
			SetLevel(tt.level)
			l := GetLevel()
			if l != tt.level {
				t.Errorf("got %d, want %d", l, tt.level)
			}
		})
	}
}

func TestPrefix(t *testing.T) {
	var logSuffix = "\x1b[0m "
	var levelTests = []struct {
		name  string
		level messageLevel
	}{
		{"\x1b[31mFATAL", fatal},
		{"\x1b[31mERROR", error},
		{"\x1b[33mWARNING", warn},
		{"\x1b[30mLOG", log},
		{"\x1b[34mINFO", info},
		{"\x1b[32mDEBUG", debug},
	}

	// Test that prefix is colored
	for _, tt := range levelTests {
		t.Run(tt.name, func(t *testing.T) {

			SetLevel(int(tt.level))
			levelPrefix := prefix(tt.level)

			// Check that we start with the color prefix
			if !strings.HasPrefix(levelPrefix, tt.name) {
				t.Errorf("got prefix %s, want %s", levelPrefix, tt.name)
			}

			// The suffix should be consistently the "off" color string
			if !strings.HasSuffix(levelPrefix, logSuffix) && tt.level != debug {
				t.Errorf("%s does not end with %s", levelPrefix, logSuffix)
			}
		})
	}
}

func TestDisableColor(t *testing.T) {
	var logSuffix = "\x1b[0m "
	var levelTests = []struct {
		name  string
		level messageLevel
	}{
		{"FATAL", fatal},
		{"ERROR", error},
		{"WARNING", warn},
		{"LOG", log},
		{"INFO", info},
		{"DEBUG", debug},
	}

	// Disable all color output, removing prefix and off suffix
	DisableColor()

	// Test that prefix is not colored
	for _, tt := range levelTests {
		t.Run(tt.name, func(t *testing.T) {

			SetLevel(int(tt.level))
			levelPrefix := prefix(tt.level)

			// Check that we don't start with the color prefix
			if !strings.HasPrefix(levelPrefix, tt.name) {
				t.Errorf("got prefix %s, want %s", levelPrefix, tt.name)
			}

			// The suffix should not contain the the "off" color string
			if strings.HasSuffix(levelPrefix, logSuffix) && tt.level != debug {
				t.Errorf("%s does ends with %s", levelPrefix, logSuffix)
			}
		})
	}
}
