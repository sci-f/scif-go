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
	"bufio"
	"io"
        "os"
        "strings"

	"github.com/sci-f/scif-go/internal/pkg/logger"
	"github.com/sci-f/scif-go/pkg/util"
)

func (client ScifClient) Setup() {
	logger.Debugf("Running Setup()")
	//ScifRecipe._exec = _exec
}


func (client ScifClient) Load(path string, apps []string, writable bool) *ScifClient {
        logger.Infof("Running Load()")

        // If the recipe is not provided (empty string) set it to be the base.
        if path == "" {
                path = Scif.Base
        }


        // Check if we have a file or a directory
        if fp, err := os.Stat(path); err == nil {

                // Case 1: It's a directory on the filesystem (scif base)
        	if fp.IsDir() {

                        // Load the filesystem and exit on error
                        if err := client.loadFilesystem(path); err != nil {
                                logger.Exitf("%s", err)
                        }

                // Case 2: It's a path to a recipe
        	} else {

                        // Load the recipe and exit on error
                        if err := client.loadRecipe(path); err != nil {
                                logger.Exitf("%s", err)
                        }
        	}

        // Otherwise, not a recipe or directory, development mode
        } else {
                logger.Warningf("No recipe or filesystem loaded, development mode.")
        }

        // TODO: after we Load, we need to update the environment
        //self.update_env(app)
        return &client
}

// loadRecipe is called on Load() if the path provided is a recipe file. It 
// should populate the Scif.config map (string|string) with various recipe
// objects
func (client ScifClient) loadRecipe(path string) error {
        logger.Infof("Calling loadRecipe, recipe %s", path)        

        // Exit quickly if file doesn't exist
        if _, err := os.Stat(path); err != nil {
                return err
        }

        // Read the file
        file, err := os.Open(path)
        defer file.Close()
        if err != nil {
                return err
        }

        // Read each line with a reader into list of lines
        var line string
        var lines []string

        reader := bufio.NewReader(file)

        for {
                line, err = reader.ReadString('\n')

                // Break when we are done
                if err != nil {
                    break
                }

                // Trim the line, remove newline, add to list
                line = strings.Trim(line, "\n")
                lines = append(lines, line)
        }

        // End of file is a successful read
        if err != io.EOF {
                return err
        }

        // We now need to populate lines into Scif.config
        section := ""
        name := ""
        var parts []string

        // Process each line
        for len(lines) > 0 {

                // Pop the first off the array
                line, lines = lines[0], lines[1:]

                // Skip comments
                if strings.HasPrefix(line, "#") {
                        continue

                // A New Section
                } else if strings.HasPrefix(line, "%") {

                        // Remove comments
                        line = strings.Split(line, "#")[0]

                        // Is there a section name?
                        parts = strings.Split(line, " ")
                        if len(parts) > 1 {
                                name = strings.Join(parts[1:]," ")          
                                logger.Debugf("Found new section name %s", name)
                        }

                        // The section is the first part, minus the %, must be lowercase
                        section = strings.Replace(parts[0], "%", "", 2)
                        section = strings.ToLower(section)
                        logger.Debugf("Found new section type %s", section)

//                config = add_section(config=config,
//                                     section=section,
//                                     name=name)

                // If we already have a section, we are adding to it
                } else if section != "" {

                        // Add the line back to parse the section to Scif.config
                        lines = util.Prepend(line, lines)
                        readSection(lines, section, name)

                }
        }


//        # Make sure app environments are sourced as first line of recipe
//        config = finish_recipe(config)   
//        
//    else:
//        bot.debug("Cannot find recipe file %s" %path)
//    return config

// TODO this should load the recipe as self.config
//             self._config = load_recipe(path)
        return nil
}

// Read a section into Scif.config, stop when we hit the next section
func readSection(lines []string, section string, name string) {
        logger.Infof("%s", Scif.config)

        // Current members of the section will be added here
        var members []string
        var nextLine string
        stripThese := []string{"applabels", "appfiles", "appenv"} 

        for {
                // If the lines are empty, break
                if len(lines) == 0 { 
                        break 
                }

                // Peek at the next line, don't remove from array
                nextLine = lines[0]
                nextLine = strings.Trim(nextLine, " ")

                // Check if the next line is a new section
                if strings.HasPrefix(nextLine, "%") {
                        break

                // Otherwise, add the nextLine to members (now remove)
                } else {
                        lines = lines[1:]

                        // If it's not a comment
                        if !strings.HasPrefix(nextLine, "#") {

                                // Strip whitespace for labels, files, environment
                                if util.Contains(section, stripThese) {
                                        nextLine = strings.Trim(nextLine, " ")
                                }
                                members = append(members, nextLine)
                        }
                }
        }

        // Add the list to the config
        if len(members) > 0 {
                if (section != "" && name != "") {
                        logger.Infof(section)
                        // Scif.config["apps"]
                }
//        if section is not None and name is not None:
//            config[global_section][name][section] = members
//        else: # section is None, is just global
//            config[global_section] = members

        }


//    return config


//def add_section(config, section, name=None, global_section="apps"):
//    ''' add section will add a section (and optionally)
//        section name to a config

//        Parameters
//        ==========
//        config: the config (dict) parsed thus far
//        section: the section type (e.g., appinstall)
//        name: an optional name, added as a level (e.g., google-drive)

//        Resulting data structure is:

//            config['registry']['apprun']
//            config[name][section]

//    '''

//    if section is None:
//        bot.error('You must define a section (e.g. %appenv) before any action.')
//        sys.exit(1)

//    if section not in sections:
//        bot.error("%s is not a valid section." %section)
//        sys.exit(1)

//    # Add the global section, if doesn't exist
//    if global_section not in config:
//        config[global_section] = OrderedDict()

//    if name is not None:        
//        if name not in config[global_section]:
//            config[global_section][name] = OrderedDict()

//        if section not in config[global_section][name]: 
//            config[global_section][name][section] = []
//            bot.debug("Adding section %s %s" %(name, section))

//    return config

}


// loadFilesystem is called if the path provided is a Scif base (directory)
func (client ScifClient) loadFilesystem(path string) error {
        logger.Debugf("path %s", path)  
// TODO this should load the filesystem as self.config
//             self._config = load_recipe(path)
        return nil
}


//def load_filesystem(base, quiet=False):
//    '''load a filesystem based on a root path, which is usually /scif

//        Parameters
//        ==========
//        base: base to load.

//        Returns
//        =======
//        config: a parsed recipe configuration for SCIF
//    '''
//    from scif.defaults import SCIF_APPS

//    if os.path.exists(SCIF_APPS):
//        apps = os.listdir(SCIF_APPS)
//        config = {'apps': {}}
//        for app in apps:
//            path = '%s/%s/scif/%s.scif' %(SCIF_APPS, app, app)
//            if os.path.exists(path):
//                recipe = load_recipe(path)
//                config['apps'][app] = recipe['apps'][app]

//        if len(config['apps']) > 0:
//            if quiet is False:
//                bot.info('Found configurations for %s scif apps' %len(config['apps']))
//                bot.info('\n'.join(list(config['apps'].keys())))
//            return config




//def finish_recipe(config, global_section='apps'):
//    '''
//       finish recipe includes final steps to add to the runtime for an app.
//       Currently, this just means adding a command to source an environment
//       before running, if appenv is defined. The Python should handle putting
//       variables in the environment, however in some cases (if the variable
//       includes an environment variable:

//          VARIABLE1=$VARIABLE2

//       It would not be properly sourced! So we add a source as the first
//       line of the runscript

//       Parameters
//       ==========
//       config: the configuation file produced by load_recipe. Assumed to have
//               a highest key of "apps" and then lookup by individual apps,
//               and then sections. Eg: config['apps']['myapp']['apprun'] 

//    '''
//    # The apps are the keys under global section "apps"
//    apps = list(config[global_section].keys())

//    for app in apps:

//        # If an apprun is present and the system supports source, do it.
//        if "appenv" in config[global_section][app]:
//            appenv = config[global_section][app]['appenv']

//            # If runscript or test is defined, add source to first line      
//            if "apptest" in config[global_section][app]:
//                apptest = config[global_section][app]['apptest']
//                config[global_section][app]['apptest'] =  appenv + apptest

//            if "apprun" in config[global_section][app]:
//                apprun = config[global_section][app]['apprun']
//                config[global_section][app]['apprun'] =  appenv + apprun

//    return config


//def read_section(config, spec, section, name, global_section='apps'):
//    '''read in a section to a list, and stop when we hit the next section
//    '''
//    members = []

//    while True:

//        if len(spec) == 0:
//            break
//        next_line = spec[0]                

//        if next_line.upper().strip().startswith("%"):
//            break
//        else:
//            new_member = spec.pop(0)
//            if not new_member.strip().startswith('#'):

//                # Strip whitespace for labels, files, environment
//                if section in ['applabels', 'appfiles', 'appenv']:
//                    new_member = new_member.strip()

//                members.append(new_member)

//    # Add the list to the config
//    if len(members) > 0:
//        if section is not None and name is not None:
//            config[global_section][name][section] = members
//        else: # section is None, is just global
//            config[global_section] = members

//    return config


//def add_section(config, section, name=None, global_section="apps"):
//    ''' add section will add a section (and optionally)
//        section name to a config

//        Parameters
//        ==========
//        config: the config (dict) parsed thus far
//        section: the section type (e.g., appinstall)
//        name: an optional name, added as a level (e.g., google-drive)

//        Resulting data structure is:

//            config['registry']['apprun']
//            config[name][section]

//    '''

//    if section is None:
//        bot.error('You must define a section (e.g. %appenv) before any action.')
//        sys.exit(1)

//    if section not in sections:
//        bot.error("%s is not a valid section." %section)
//        sys.exit(1)

//    # Add the global section, if doesn't exist
//    if global_section not in config:
//        config[global_section] = OrderedDict()

//    if name is not None:        
//        if name not in config[global_section]:
//            config[global_section][name] = OrderedDict()

//        if section not in config[global_section][name]: 
//            config[global_section][name][section] = []
//            bot.debug("Adding section %s %s" %(name, section))

//    return config
