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

package main

import (
	"errors"
	"fmt"
	"os"
	//"os/user"
	//"path"
	"strings"
	"text/template"

	"github.com/sci-f/scif-go/cmd/scif/docs"
	"github.com/sci-f/scif-go/pkg/version" // version
	"github.com/spf13/cobra"
	//"github.com/spf13/pflag"
	// each import here should be from /internal/pkg/<folder>
	"github.com/sci-f/scif-go/internal/pkg/logger"
)

// Global variables
var (
	base     string
	debug    bool
	nocolor  bool
	quiet    bool
	silent   bool
	readonly bool
)

// COMMANDS ....................................................................

func init() {
	ScifCmd.Flags().SetInterspersed(false)
	ScifCmd.PersistentFlags().SetInterspersed(false)

	templateFuncs := template.FuncMap{
		"TraverseParentsUses": TraverseParentsUses,
	}
	cobra.AddTemplateFuncs(templateFuncs)

	ScifCmd.SetHelpTemplate(docs.HelpTemplate)
	ScifCmd.SetUsageTemplate(docs.UseTemplate)

	// Set a custom version template string
	vt := fmt.Sprintf("scif version {{printf \"%%s\" .Version}}\n", version.Version)
	ScifCmd.SetVersionTemplate(vt)

	ScifCmd.Flags().BoolVarP(&debug, "debug", "d", false, "print debugging information (highest verbosity)")
	ScifCmd.Flags().BoolVar(&nocolor, "nocolor", false, "print without color output (default False)")
	ScifCmd.Flags().BoolVarP(&silent, "silent", "s", false, "only print errors")
	ScifCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "suppress normal output")
	ScifCmd.Flags().BoolVarP(&readonly, "readonly", "r", false, "use scif without writable")

	VersionCmd.Flags().SetInterspersed(false)
	ScifCmd.AddCommand(VersionCmd)

}

// ScifCmd is the command base, when the user calls "scif" without subcommands
var ScifCmd = &cobra.Command{
	TraverseChildren:      true,
	DisableFlagsInUseLine: true,
	PersistentPreRun:      runSetup,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("Invalid command")
	},

	Use:           docs.ScifUse,
	Version:       version.Version,
	Short:         docs.ScifShort,
	Long:          docs.ScifLong,
	Example:       docs.ScifExample,
	SilenceErrors: true,
	SilenceUsage:  true,
}

// TraverseParentsUses walks the parent commands and outputs a properly formatted use string
func TraverseParentsUses(cmd *cobra.Command) string {
	if cmd.HasParent() {
		return TraverseParentsUses(cmd.Parent()) + cmd.Use + " "
	}

	return cmd.Use + " "
}

// LOGGING .....................................................................

// setLoggerLevel: set the logging level, with default 1 (info0
func setLoggerLevel(cmd *cobra.Command, args []string) {
	var level int

	if debug {
		level = 5
	} else if quiet {
		level = -1
	} else if silent {
		level = -3
	} else {
		level = 1
	}
	logger.SetLevel(level)
}

// setLoggerColor allows the user to disable the color
func setLoggerColor(cmd *cobra.Command, args []string) {
	if nocolor {
		logger.DisableColor()
	}
}

func runSetup(cmd *cobra.Command, args []string) {
	setLoggerLevel(cmd, args)
	setLoggerColor(cmd, args)
}

// ENTRYPOINT ..................................................................
func main() {
	ExecuteScif()
}

// ExecuteScif adds all child commands to the root command and sets flags
func ExecuteScif() {
	if cmd, err := ScifCmd.ExecuteC(); err != nil {
		if str := err.Error(); strings.Contains(str, "unknown flag: ") {
			flag := strings.TrimPrefix(str, "unknown flag: ")
			ScifCmd.Printf("Invalid flag %q for command %q.\n\nOptions:\n\n%s\n",
				flag,
				cmd.Name(),
				// Return flag usage wrapped (0 means no columns)
				cmd.Flags().FlagUsagesWrapped(0))
		} else {
			ScifCmd.Println(cmd.UsageString())
		}
		ScifCmd.Printf("Run '%s --help' for more detailed usage information.\n",
			cmd.CommandPath())
		os.Exit(1)
	}
}

// VERSION .....................................................................

// VersionCmd displays installed scif version
var VersionCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.Version)
	},

	Use:   "version",
	Short: "Show the Scientific Filesystem Version",
}
