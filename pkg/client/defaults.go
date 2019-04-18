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
	"os"
	"strings"

	"github.com/sci-f/scif-go/internal/pkg/logger"
)

// getenv will return an environment variable, or return a default
func getenv(key string, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	logger.Debugf("Key %s found as %s", key, value)
	return value
}

// getBoolEnv returns a boolean environment variable
func getBoolEnv(key string, fallback bool) bool {
	value := os.Getenv(key)

	// If the value isn't set, return False
	if len(value) == 0 {
		return fallback
	}

	// Ensure lowercase
	value = strings.ToLower(value)

	// Strings that are valid to indicate a setting of "True"
	list := [5]string{"yes", "true", "t", "1", "y"}
	for _, yes := range list {
		if value == yes {
			return true
		}
	}
	return false
}

//scifAllowAppend :=
// TODO: need to add SCIF_APPS getenv_namespace(namespace="SCIF_APP")
// TODO: getStringDefault will return empty string (not nil) if not found
// scifDefaults are grabbed from the environment, e.g., SCIF_BASE

// getStringDefault returns the default for a string, or empty string
// if you need to add a new default from the environment, define it here
// and add it to client.go under ScifRecipe and NewScifRecipe (the init)
func getStringDefault(key string) string {
	defaults := map[string]string{

		"BASE":        "/scif",
		"DATA":        "/scif/data",
		"APPS":        "/scif/apps",
		"SHELL":       "/bin/bash",
		"ENTRYPOINT":  "/bin/bash",
		"ENTRYFOLDER": "",
	}

	if value, ok := defaults[key]; ok {
		return value
	}
	return ""
}

// TODO: need to lookup how to correctly pass error if not there.
// getBoolDefault returns the default for a bool, or false
func getBoolDefault(key string) bool {

	defaults := map[string]bool{
		"ALLOW_APPEND_PATHS": true,
	}

	if value, ok := defaults[key]; ok {
		return value
	}
	return false
}

// An array of paths to append
// return all environment variables in a particular namespace, such as for
// SCIF apps based on starting with SCIF_APPS.
func getenvNamespace(prefix string) []string {

	var namespace []string

	for _, env := range os.Environ() {
		// pair[0] is the key, pair[1] is value
		pair := strings.Split(env, "=")

		// If the key starts with prefix, add to namespace list
		if strings.HasPrefix(pair[0], prefix) {
			namespace = append(namespace, pair[1])
		}
	}

	return namespace
}
