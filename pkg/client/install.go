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

package client

import (
        "os"

	"github.com/sci-f/scif-go/internal/pkg/logger"
	"github.com/sci-f/scif-go/pkg/util"
	// jRes, err := util.ParseErrorBody(resp.Body)
)

// Install an app for a scientific filesystem
// install recipes to a base. We assume this is the root of a system
// or container, and will write the /scif directory on top of it.
// If an app name is provided, install that app if it is found 
// in the config. This function goes through all steps to:
//
// 1. Install base folders to base, creating a folder for each app
// 2. Install one or more apps to it, the config is already loaded

func Install(recipe string, apps []string, writable bool) (err error) {

	logger.Debugf("Installing recipe %s", recipe)

        // Ensure that recipe exists
        if _, err := os.Stat(recipe); os.IsNotExist(err) {
                logger.Exitf("Recipe %s does not exist.", recipe)
        }

        // Ensure we have writable if asking for it
        if (writable && !util.HasWriteAccess(Scif.Base)) {
                logger.Exitf("No write access to %s", Scif.Base)
        }

        // Create the client
	cli := ScifClient{}

        // install Base folders
	cli.installBase()
        cli.installApps(apps)

	return err
}


// Install Helper Functions
// these functions are added to the ScifClient struct base, and have access
// to other variables via the (initialized) Scif.<varname>)
// .............................................................................

// installBase is a private function to install the base, apps, and data folder
func (client ScifClient) installBase() {
	logger.Infof("Installing base to %s", Scif.Base)

        // Create the base, apps folder, and data folders
        folders := []string{Scif.Base, Scif.Apps, Scif.Data}

        // Exit on any kind of error
        for _, folder := range folders {
                if err := os.MkdirAll(folder, os.ModePerm); err != nil {
                        logger.Exitf("%s", err)
                }
        }
}

// installApps installs one or more apps to the base, apps is a list of apps.
func (client ScifClient) installApps(apps []string) {

        // If no apps defined, get those found at base        

        // Loop through apps to install
        for _, app := range apps {

                logger.Infof(app)
        } 
}
//def install_apps(self, apps=None):
//    '''install one or more apps to the base. If app is defined, only
//       install app specified. Otherwise, install all found in config.
//    '''
//    if apps in [None, '']:
//        apps = self.apps()

//    if not isinstance(apps, list):
//        apps = [apps]

//    if len(apps) == 0:
//        bot.warning('No apps to install. Load a recipe or base with .load()')

//    for app in apps:

//        # We must have the app defined in the config
//        if app not in self._config['apps']:
//            bot.error('Cannot find app %s in config.' %app)
//            sys.exit(1)

//        # Make directories
//        settings = self._init_app(app)

//        # Get the app configuration
//        config = self.app(app)

//        # Get the app environment and export for install
//        self.get_appenv(app, isolated=False, update=True)
//        self.export_env(ps1=False)

//        # Handle environment, runscript, labels
//        self._install_runscript(app, settings, config)
//        self._install_environment(app, settings, config)
//        self._install_help(app, settings, config)
//        self._install_labels(app, settings, config)
//        self._install_files(app, settings, config)
//        self._install_commands(app, settings, config)
//        self._install_recipe(app, settings, config)
//        self._install_test(app, settings, config)

//        # After we install, in case interactive, deactivate last app
//        self.deactivate(app)





//def install(self, app=None):
//    '''install recipes to a base. We assume this is the root of a system
//       or container, and will write the /scif directory on top of it.
//       If an app name is provided, install that app if it is found 
//       in the config. This function goes through all steps to:

//       1. Install base folders to base, creating a folder for each app
//       2. Install one or more apps to it, the config is already loaded
//    '''

//    self._install_base()             # Generate the folder structure
//    self._install_apps(app)          # App install


//def init_app(self, app):
//    '''initialize an app, meaning adding the metadata folder, bin, and 
//       lib to it. The app is created at the base
//    '''
//    settings = self.get_appenv_lookup(app)[app]

//    # Create base directories for metadata
//    for folder in ['appmeta', 'appbin', 'applib', 'appdata']:
//        mkdir_p(settings[folder])
//    return settings



//def install_labels(self, app, settings, config):
//    '''install labels will add labels to the app labelfile

//       Parameters
//       ==========
//       app should be the name of the app, for lookup in config['apps']
//       settings: the output of _init_app(), a dictionary of environment vars
//       config: should be the config for the app obtained with self.app(app)

//    '''
//    lookup = dict()
//    if "applabels" in config:
//        labels = config['applabels']
//        bot.level
//        bot.info('+ ' + 'applabels '.ljust(5) + app)
//        for line in labels:
//            label, value = get_parts(line, default='')
//            lookup[label] = value 
//        write_json(lookup, settings['applabels'])
//    return lookup


//def install_files(self, app, settings, config):
//    '''install files will add files (or directories) to a destination.
//       If none specified, they are placed in the app base

//       Parameters
//       ==========
//       app should be the name of the app, for lookup in config['apps']
//       settings: the output of _init_app(), a dictionary of environment vars
//       config: should be the config for the app obtained with self.app(app)

//    '''
//    if "appfiles" in config:
//        files = config['appfiles']
//        bot.info('+ ' + 'appfiles '.ljust(5) + app)

//        for pair in files:
//        
//            # Step 1: determine source and destination
//            src, dest = get_parts(pair, default=settings['approot'])

//            # Step 2: copy source to destination
//            cmd = ['cp']

//            if os.path.isdir(src):
//                cmd.append('-R')
//            elif os.path.exists(src):
//                cmd = cmd + [src, dest]
//                result = self._run_command(cmd)
//            else:    
//                bot.warning('%s does not exist, skipping.' %src)



//def install_commands(self, app, settings, config):
//    '''install will finally, issue commands to install the app.

//       Parameters
//       ==========
//       app should be the name of the app, for lookup in config['apps']
//       settings: the output of _init_app(), a dictionary of environment vars
//       config: should be the config for the app obtained with self.app(app)

//    '''
//    if "appinstall" in config:

//        # Change directory so the APP is $PWD
//        pwd = os.getcwd()
//        os.chdir(settings['approot'])
//        
//        # issue install commands
//        cmd = '\n'.join(config['appinstall'])
//        bot.info('+ ' + 'appinstall '.ljust(5) + app)
//        os.system(cmd)

//        # Go back to previous location
//        os.chdir(pwd)


//def install_recipe(self, app, settings, config):
//    '''Write the initial recipe for the app to its metadata folder.

//       Parameters
//       ==========
//       app should be the name of the app, for lookup in config['apps']
//       settings: the output of _init_app(), a dictionary of environment vars
//       config: should be the config for the app obtained with self.app(app)

//    '''
//    recipe_file = settings['apprecipe']
//    recipe = '' 

//    for section_name, section_content in config.items():
//        content = '\n'.join(section_content)
//        header = '%' + section_name
//        recipe += '%s %s\n%s\n' %(header, app, content)

//    write_file(recipe_file, recipe)
//    return recipe


//# Scripts

//def install_script(self, section, app, settings, config, executable=False):
//    '''a general function used by install_runscript, install_help, and
//       install_environment to write a script to a file from a config setting
//       section

//       Parameters
//       ==========
//       section: should be the name of the section in the config (e.g., apprun)
//       app should be the name of the app, for lookup in config['apps']
//       settings: the output of _init_app(), a dictionary of environment vars
//       config: should be the config for the app obtained with self.app(app)
//       executable: if the file is written, make it executable (defaults False)

//    '''
//    if section in config:
//        content = '\n'.join(config[section])
//        bot.info('+ ' + section + ' '.ljust(5) + app)
//        write_file(settings[section], content)

//        # Should we make the script executable (checks for exists)
//        if executable is True:
//            make_executable(settings[section])


//def install_runscript(self, app, settings, config, executable=True):
//    '''install runscript will prepare the runscript for an app.
//       the parameters are shared by _install_script
//    '''
//    return self._install_script('apprun', app, settings, config, executable)

//            
//def install_environment(self, app, settings, config):
//    '''install will run the content to export environment variables, if defined
//       the parameters are shared by _install_script
//    '''
//    return self._install_script('appenv', app, settings, config)


//def install_help(self, app, settings, config):
//    '''install will write the help section, if defined.
//       the parameters are shared by _install_script
//    '''
//    return self._install_script('apphelp', app, settings, config)


//def install_test(self, app, settings, config, executable=True):
//    '''install test will prepare a test script for an app.
//       the parameters are shared by _install_script
//    '''
//    return self._install_script('apptest', app, settings, config, executable)
