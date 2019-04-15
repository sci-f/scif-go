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

package docs

// For each subcommand for the client, you should write a section:
//   - include a <name>Use, <name>Short, and <name>Long
//   - if relevant, include a <name>Example.

const (

	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// main scif command
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

	ScifUse   string = `scif [global options...]`
	ScifShort string = `Scientific Filesystem for modular internal organization of containers.`
	ScifLong  string = `
        The Scientific Filesystem (SCIF) provides an internal organization and interaction
        specification to allow for modular applications to coexist within the same
        reproducible container. You typically install from a recipe.scif and then interact
        with the file system via the scif client.`
	ScifExample string = `

        $ scif help <command> [<subcommand>]
        $ scif help install
        $ scif help (run|inspect|exec)`

	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// shell
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

	ShellUse   string = `shell [-h] [app]`
	ShellShort string = `Shell into a Scientific Filesystem or a specific app.`
	ShellLong  string = `
        positional arguments:
          app         app shell to, defaults to SCIF base if not set.

        optional arguments:
          -h, --help  show this help message and exit`
	ShellExample string = `

        $ scif shell
        $ scif shell <app>`

	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// preview
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

	PreviewUse   string = `preview [-h] [recipe [recipe ...]]`
	PreviewShort string = `Preview the operations to be done for an install without doing it.`
	PreviewLong  string = `  
        positional arguments:
          recipe      recipe file for the filesystem

        optional arguments:
          -h, --help  show this help message and exit`
	PreviewExample string = `

        $ scif preview <recipe>`

	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// install
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	InstallUse   string = `install [-h] [recipe [recipe ...]]`
	InstallShort string = `Install a recipe for a Scientific Filesystem`
	InstallLong  string = `
        positional arguments:
          recipe      recipe file for the filesystem

        optional arguments:
          -h, --help  show this help message and exit`
	InstallExample string = `

        $ scif install <recipe>`

	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// inspect
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	InspecttUse  string = `inspect [-h] [attributes [attributes ...]]`
	InspectShort string = `inspect attributes for a scif application`
	InpectLong   string = `
        positional arguments:
          attributes  attribute to inspect (runscript|r), (environment|e), (labels|l),
                      or (all|a) (default)

        optional arguments:
          -h, --help  show this help message and exit`
	InspectExample string = `

        $ scif inspect <app>`

	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// run
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	RunUse   string = `run [-h] [cmd [cmd ...]]`
	RunShort string = `Run a Scientific Filesystem application.`
	RunLong  string = `
        positional arguments:
          cmd         app and optional arguments to target for the entry

        optional arguments:
          -h, --help  show this help message and exit`
	RunExample string = `

        $ scif run <app>
        $ scif run <app> [args]`

	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// test
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	TestUse   string = `test [-h] [cmd [cmd ...]]`
	TestShort string = `Test a Scientific Filesystem application.`
	TestLong  string = `
        positional arguments:
          cmd         app and optional arguments to target for the entry

        optional arguments:
          -h, --help  show this help message and exit`
	TestExample string = `

        $ scif test
        $ scif test <app>`

	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// apps
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	AppsUse   string = `scif apps [-h] [-l]`
	AppsShort string = `list Scientific Filesystem Applications installed`
	AppsLong  string = `
        optional arguments:
          -h, --help  show this help message and exit
          -l          show long listing, including paths.`
	AppsExample string = `

        $ scif apps
        $ scif apps -l`

	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// dump
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	DumpUse   string = `dump [-h]`
	DumpShort string = `export the configuration for a Scientific Filesystem.`
	DumpLong  string = `
        optional arguments:
          -h, --help  show this help message and exit`
	DumpExample string = `

        $ scif dump`

	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// exec
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	ExecUse   string = `scif exec [-h] [cmd [cmd ...]]`
	ExecShort string = `execute a command to a Scientific Filesystem`
	ExecLong  string = `
        positional arguments:
          cmd         app and command to execute. Eg, exec appname echo $SCIF_APPNAME

        optional arguments:
          -h, --help  show this help message and exit`
	ExecExample string = `

        $ scif exec <app> [cmd]
        $ scif exec appname echo "Hello?"`
)
