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

package docs

// HelpTemplate is the template exported to be used in other packages
const (
	HelpTemplate string = `{{.Short}}

Usage:
  {{.UseLine}}

Description:{{.Long}}{{if .HasAvailableLocalFlags}}

Options:
{{.LocalFlags.FlagUsagesWrapped 80 | trimTrailingWhitespaces}}
{{end}}{{if .HasAvailableInheritedFlags}}

Global Options:
{{.InheritedFlags.FlagUsagesWrapped 80 | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableSubCommands}}
Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasExample}}

Examples:{{.Example}}{{end}}


For additional help or support, please post an issue at https://www.github.com/sci-f/scif-go
`

	UseTemplate string = `Usage:
  {{TraverseParentsUses . | trimTrailingWhitespaces}}{{if .HasAvailableSubCommands}} <command>

Available Commands:{{range .Commands}}{{if .IsAvailableCommand}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}
`
)
