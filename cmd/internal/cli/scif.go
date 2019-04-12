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

package cli

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"path"
	"strings"
	"text/template"

	"github.com/sci-f/scif-go/docs"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
        "github.com/sci-f/scif-go/pkg" // version
	// each import here should be from /internal/pkg/<folder>
	//"github.com/sci-f/scif-go/internal/pkg/buildcfg"
	"github.com/sci-f/scif-go/internal/pkg/logger"
)

// Global variables
var (
	debug    bool
	nocolor  bool
	quiet    bool
	silent   bool
	writable bool
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
	vt := fmt.Sprintf("%s version {{printf \"%%s\" .Version}}\n", scif.Version)
	ScifCmd.SetVersionTemplate(vt)

	ScifCmd.Flags().BoolVarP(&debug, "debug", "d", false, "print debugging information (highest verbosity)")
	ScifCmd.Flags().BoolVar(&nocolor, "nocolor", false, "print without color output (default False)")
	ScifCmd.Flags().BoolVarP(&silent, "silent", "s", false, "only print errors")
	ScifCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "suppress normal output")
	ScifCmd.Flags().BoolVarP(&writable, "writable", "w", false, "use scif with writable")

	VersionCmd.Flags().SetInterspersed(false)
	ScifCmd.AddCommand(VersionCmd)

}

// ScifCmd is the command base, when the user calls "scif" without subcommands
var ScifCmd = &cobra.Command{
	TraverseChildren:      true,
	DisableFlagsInUseLine: true,
	PersistentPreRun:      persistentPreRun,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("Invalid command")
	},

	Use:           docs.ScifUse,
	Version:       scif.Version,
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


func persistentPreRun(cmd *cobra.Command, args []string) {
	setLoggerLevel(cmd, args)
	setLoggerColor(cmd, args)
	updateFlagsFromEnv(cmd)
}


// ENTRYPOINT ..................................................................

// ExecuteScif adds all child commands to the root command and sets flags
func ExecuteScif() {
	if cmd, err := ScifCmd.Execute(); err != nil {
		if str := err.Error(); strings.Contains(str, "unknown flag: ") {
			flag := strings.TrimPrefix(str, "unknown flag: ")
			ScifCmd.Printf("Invalid flag %q for command %q.\n\nOptions:\n\n%s\n",
				flag,
				cmd.Name(),
				cmd.Flags().FlagUsagesWrapped(getColumns()))
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
var VersionCmd = &cobra.Command {
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(scif.Version)
	},

	Use:   "version",
	Short: "Show the Scientific Filesystem Version",
}
