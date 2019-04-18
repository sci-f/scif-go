# Usage

## Install

You should clone the library into your $GOPATH For example:

```bash
$ mkdir -p $GOPATH/src/github.com/sci-f
$ cd $GOPATH/src/github.com/sci-f
$ git clone https://www.github.com/sci-f/scif-go
```

Next, install dependenies, and build.

```bash
$ make deps
$ make build
```

If you aren't familiar with the organizational logic, see the [organization](organization.md) docs.
The "scif" executable will be built in the "bin" folder of the repository, so you
can interact with it like:

```bash
$ bin/scif --help
```

## Install a Recipe

Export the base to somewhere like /tmp/scif so you don't need to use sudo to write (the default is /scif).

```bash
export SCIF_BASE=/tmp/scif
```

Preview the recipe [hello-world.scif](../hello-world.scif) if you like (this produces verbose output)

```bash
$ bin/scif preview hello-world.scif
...
```

You can also just look at the file - the sections are pretty intuitive!

 - `%appinstall <name>` is a set of commands, run in the app folder, to install the application
 - `%apprun <name>` is the runscript, or what gets executed with "scif run <name>"
 - `%applabels <name>` should be key value pairs of metadata, separated by =
 - `%appenv <name>` Is a little script that will be sourced for the environment.
 - `%appfiles <name>` A list of source destination files to add to the app folder
 - `%apptest <name>` A script to run to test the app

When you are ready, run the install:

```bash
$ bin/scif install hello-world.scif 
Key SCIF_BASE found as /tmp/scif
INFO:    Installing base to /tmp/scif
INFO:    Installing app hello-world-echo
INFO:    Installing app hello-world-script
INFO:    Installing app hello-custom
INFO:    Installing app hello-world-env
```

And take a look at /tmp/scif to see the organization and resulting files!

```bash
$ tree /tmp/scif
/tmp/scif
├── apps
│   ├── hello-custom
│   │   ├── bin
│   │   ├── lib
│   │   └── scif
│   │       ├── hello-custom.scif
│   │       └── runscript
│   ├── hello-world-echo
│   │   ├── bin
│   │   ├── lib
│   │   └── scif
│   │       ├── environment.sh
│   │       ├── hello-world-echo.scif
│   │       └── runscript
│   ├── hello-world-env
│   │   ├── bin
│   │   ├── lib
│   │   └── scif
│   │       ├── environment.sh
│   │       ├── hello-world-env.scif
│   │       └── runscript.help
│   └── hello-world-script
│       ├── bin
│       │   └── hello-world.sh
│       ├── lib
│       └── scif
│           ├── environment.sh
│           ├── hello-world-script.scif
│           ├── runscript
│           └── test.sh
└── data
    ├── hello-custom
    ├── hello-world-echo
    ├── hello-world-env
    └── hello-world-script

22 directories, 13 files
```

Notice that each app folder, under "apps" is named according to the name you gave it.
Each has a "bin" and a "lib" folder that you can add files to. Anything in "bin" will
be added to the path when the app is run, and anything in "lib" will be added to LD_LIBRARY_PATH.
Apps that had an `apptest` section have a test.sh file, and apps with an entrypoint have a
runscript. These files are in a metadata folder that is also called "scif." The data folder
also has a folder for each app installed.

For details on writing recipes, the environment, and other information about the
Scientific Fileystem see [sci-f.github.io](https://sci-f.github.io).
