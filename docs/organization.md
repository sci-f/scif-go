# Organization

This is how I decided to organize files and folders. I wanted everything to be located
where it would intuitively be looked for.

## Where is the version?

The version is kept in [pkg/version/version.go](pkg/version/version.go) and then piped into the Makefile and 
imported into the client to print. This isn't the only way to do it, but seemed reasonable.


## Where is the scif client (command) entrypoint?

[cmd/scif](cmd/scif) is the entrypoint command for the application. The file [main.go](cmd/scif/main.go) has the main client, and the other *.go files in the folder correspond with subcommands (e.g., "run.go"). Flags that are global are in flags.go, and environment.go has environment parsing. Generally, you can find what you are looking for based on the naming, despite the fact that all these files are in shared "package main" and get compiled together.

## Where are the client functions?

[pkg/client](pkg/client) is logically a package for the client functions. These are the functions that are called from the entrypoint client, the base of which is defined in [pkg/client/client.go](pkg/client/client.go). The execution first creates a struct that is going to hold client variables:

```go
type ScifClient struct {
	Base        string   // /scif is the overall base
	Data        string   // <Base>/data is the data base
	Apps        string   // <Base>/apps is the apps base
	ShellCmd    string   // default shell
...
	activeApp   string // the active app (if one is defined)
}
```

and then we create an initialization function that is going to help define some of those variables, and
do other setup tasks. Notice how it returns an instantiated version of ScifClient:

```go
func NewScifClient() *ScifClient {

	base := getenv("SCIF_BASE", getStringDefault("BASE"))
	scifApps := getenvNamespace("SCIF_APP")

...

	// Instantiate the client
	client := &ScifClient{Base: base,
		Data:        data,
		Apps:        apps,
		ShellCmd:    shell,
		EntryPoint:  entrylist,
		EntryFolder: entryfolder,
		allowAppend: allowAppend,
		appendPaths: scifAppendPaths,
		scifApps:    scifApps}

	// Setup includes loading a scif, if found at base
	client.Setup()

	return client
}
```


And thus, this object (?) is available to the user (with variables defined) as "Scif":

```go
// provide client to user as "Scif"
var Scif ScifClient = *NewScifClient()
```

## How do files relate between the two?

Thus, when we call a function in [pkg/client](pkg/client) from an entrypoint in [cmd/scif](cmd/scif)
we call the `<package>/<function>` directly, and it's usually the case they are found in matching files (e.g., install.go
and install.go in each directory). The example below shows calling Install in the client package after 
parsing input arguments for a recipe, additional arguments, and a boolean:

```go
client.Install(recipe, args, !readonly)
```

## How do we instantiate the client?

Then in the package "client" install.go we might check that the recipe exist, and create an instance of the 
ScifRecipe struct that has all the helper functions attached. If we want to load a recipe that is provided, 
we create the client like this:

```go
// Create the client, load the recipe
cli := ScifClient{}.Load(recipe, apps, writable)
```

Otherwise we don't call load (and could call it later, if desired):

```go
// Create the client
cli := ScifClient{}
```

Either way, after we have loaded, we can further call functions that are owned by the client.

```go
// install Base folders
cli.installBase()
cli.installApps(apps)
```

## How do we add functions to the client?

We add functions to the ScifRecipe like this:

```go
func (client ScifClient) Execute() {

	logger.Debugf("Execute() here")
	fmt.Println("The base is at %s", Scif.Base)
}
```

And notice how we reference the variables that have been initialized via Scif.Base.
