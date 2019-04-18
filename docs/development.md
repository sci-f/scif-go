# Development

## Setup

### 1. Clone from GitHub

When you are first developing the library, you should clone it into your $GOPATH.
For example:

```bash
$ mkdir -p $GOPATH/src/github.com/sci-f
$ cd $GOPATH/src/github.com/sci-f
$ git clone https://www.github.com/sci-f/scif-go
```

### 2. Install and Build

Next, install dependenies, and try a build.

```bash
$ make deps
$ make build
LDFLAGS: -ldflags -X=main.Version=0.0.1.rc -X=main.Build=1b5937ef84554706663938316abfe2cf223ef530
GOFILES: ./internal/pkg/logger/logger.go
./pkg/version/version.go
./pkg/util/entrypoint.go
./pkg/util/agent/agent.go
./pkg/util/io.go
./pkg/util/strings.go
./pkg/client/install.go
./pkg/client/apps.go
./pkg/client/client.go
./pkg/client/environment.go
./pkg/client/run.go
./pkg/client/export.go
./pkg/client/setup.go
./pkg/client/preview.go
./pkg/client/defaults.go
./cmd/scif/main.go
./cmd/scif/flags.go
./cmd/scif/install.go
./cmd/scif/environment.go
./cmd/scif/run.go
./cmd/scif/docs/content.go
./cmd/scif/docs/templates.go
./cmd/scif/preview.go
```

If you aren't familiar with the organizational logic, see the [organization](organization.md) docs.
The "scif" executable will be built in the "bin" folder of the repository, so you
can interact with it like:

```bash
$ bin/scif --help
```

At this point, you can make a choice to install and develop locally, or to
install and develop using a Docker container. 

## Locally

If you want to develop locally, it's highly recommended to export the SCIF_BASE
to be somewhere *other* than the default at `/scif`. The reason is because using
the default would require sudo, and if you are developing (and writing files)
you generally don't want to muck anything up :) So it is recommended
to define a new base:

```bash
export SCIF_BASE=/tmp/scif
```

You can then test (what would be created) with a preview. Note that a dummy
recipe, [hello-world.scif](../hello-world.scif) is provided in the root of the 
repository. Try previewing a single app from the recipe to see the install base:

```bash
$ bin/scif preview hello-world.scif hello-custom
DEBUG                  getenv()                      Key SCIF_BASE found as /tmp/scif
DEBUG                  getenv()                      Key SCIF_BASE found as /tmp/scif
DEBUG                  getenv()                      Key SCIF_BASE found as /tmp/scif
INFO:    [base] /tmp/scif
INFO:    [apps] /tmp/scif
INFO:    [data] /tmp/scif
INFO:    

hello-custom
...
```

Now here is how your development workflow can work. You can safely do iterations of
building, running the application (with the base in /tmp/scif) and then printing to
see changes.

```bash
$ make build
$ bin/scif <command> <args>...
```

## Docker

As an alternative, you can build a Dockerized version, and then bind the
folder with code you are working on, and rebuild in the container. Note that
you want to use Dockerfile.dev (a non multistage build) to maintain the dependencies
in the container for use again. I haven't
done this yet, but generally you'd build like this:

```bash
$ docker build -f Dockerfile.dev -t vanessa/scif-go .
```

And then bind:


```bash
$ docker run -it --entrypoint sh -v $PWD:/go/src/github.com/sci-f/scif-go vanessa/scif-go
```

And then cd to that location, and issue the same make commands to build and install.
Be careful that you use the binary in bin/scif (and not the one moved to /usr/local.
