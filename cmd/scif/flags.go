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

package main

import (
	//"log"
	//"os"

	"github.com/spf13/pflag"
)

// https://godoc.org/github.com/spf13/pflag
// Most of scif include positional arguments, with a few booleans here
var (
	AppsLongList     bool // longlist apps (with paths)
	InspectRunscript bool
	InspectEnv       bool
	InspectLabels    bool
	InspectAll       bool // default for inspect
)

// boolFlags instantiates the flag set, to be used by subcommands
var boolFlags = pflag.NewFlagSet("BoolFlags", pflag.ExitOnError)

func init() {
	initBoolFlags()
}

// initBoolFlags initializes the few boolean arguments for scif
func initBoolFlags() {

	// -l|--longlist
	boolFlags.BoolVarP(&AppsLongList, "longlist", "L", false, "show long listing, including paths")
	boolFlags.SetAnnotation("longlist", "envkey", []string{"LONGLIST"})

	// -r|--runscript
	boolFlags.BoolVarP(&InspectRunscript, "runscript", "r", false, "inspect the runscript for one or more scientific filesystem applications.")
	boolFlags.SetAnnotation("runscript", "envkey", []string{"RUNSCRIPT"})

	// -e|--environment
	boolFlags.BoolVarP(&InspectEnv, "environment", "e", false, "inspect the environment for one or more scientific filesystem applications.")
	boolFlags.SetAnnotation("environment", "envkey", []string{"ENVIRONMENT"})

	// -l|--labels
	boolFlags.BoolVarP(&InspectEnv, "labels", "l", false, "inspect the labels for one or more scientific filesystem applications.")
	boolFlags.SetAnnotation("labels", "envkey", []string{"LABELS"})

	// -a|--all
	boolFlags.BoolVarP(&InspectEnv, "all", "a", false, "inspect all attributes for one or more scientific filesystem applications.")
	boolFlags.SetAnnotation("all", "envkey", []string{"ALL"})

}
