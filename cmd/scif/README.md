# Scif Client

## How is this organized?

 - [scif.go](scif.go) is the main file that serves the init process for the client.

Then each corresponding file is a subcommand (e.g., install.go). You can find
the flags for each subcommand in their file - I've seen a lot of shared flags.go /
environment.go files, and I find this organization just confusing. If I have to
spend time explaining it, it's too much.

## How does the environment come in?

Largely, most clients will (first preference) use command line flags, and then
fall back to the environment (if set). The actual parsing of the environment
is usually done in the client package (e.g., for [cmd/scif/install.go](install.go), the
client package would initialize the environment in [pkg/client/defaults.go](../../pkg/client/defaults.go)
and then run logic to decide the final variable in [pkg/client/install.go](../../pkg/client/install.go)

## What are the subcommands?

Each of the subcommands has its own file, within it we define the sub command
group and add any required flags from it. The commands, generally, are:

```
Shell:
  app: shell to act with scientific filesystem (not required, defaults to scif base if not set)

Preview: preview changes to a filesystem
  recipe: recipe file for the filesystem. If user provides more than one argument they are apps

Help:
  app: app(s) to print help for.

Install: install a recipe on the filesystem
  recipe: recipe for the filesystem

Inspect: inspect an attribute for a scientific filesystem installation
  attributes: attributes to inspect (runscript|r), (environment|e), (labels|l), (all|a) <- default

Run: entrypoint to run a scientific filesystem
  cmd: app and optional arguments to target for the entry

Test: entrypoint to test an app in a scientific filesystem
  cmd: app and optional arguments to target for the entry

Apps: list apps installed
 -l longlist, show long listing, including paths

Dump: dump a recipe
  
Execute: execute a command to the scientific filesystem
  cmd: app and command to execute (e.g., exec <appname> echo "HELLO"
```

If you need any more documentation or notes, please ask.
