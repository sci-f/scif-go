# SCI-F GoLang

This is the GoLang implementation of the Scientific Filesystem, and the first full GoLang package
that I'm developing. 

**under development**

## Development

Here is how I figured this out, because I really didn't have a sense of what I was doing.

 1. I started with an [empty project template](https://github.com/golang-standards/project-layout).
 2. I read through the (original) README.md for a high level overview of the structure, and rewrote the sections (below) to further develop my own understanding.
 3. I read through [this post](https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2) carefully to understand the repository structure, and what should go in each folder. For each section, I would inspect my cloned repository, and look at the README.md in the folder of inspection. You will find examples! It's essential to read about the folder's purpose, and then look at a lot of examples to see it in action. Only stop looking at examples when you "get it."
 4. I then started with the cmd folder, and slowly started writing (skeleton / dummy) implementations for what I thought would work. If there was an example that I liked that I wanted to remember to do, I added it to the TODO list below. This came down to:
   a. First noticing that the `cmd/internal/cli/scif.go` could import documentation templates, and writing those.

### TODO

 - there should be programmatic docs in a (top level) docs folder that is imported in cmd/internal/cli/scif.go (the help strings, etc.)
 - create examples in examples folder, and testing on circle
 - Add badges: [Go Report Card](https://goreportcard.com/), [GoDoc](http://godoc.org), release, [![Go Report Card](https://goreportcard.com/badge/github.com/golang-standards/project-layout?style=flat-square)](https://goreportcard.com/report/github.com/golang-standards/project-layout) [![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/golang-standards/project-layout) [![Release](https://img.shields.io/github/release/golang-standards/project-layout.svg?style=flat-square)](https://github.com/golang-standards/project-layout/releases/latest)
 - internal/pkg/logger should have the logger

## Variables

Version. The version is kept in [pkg/version.go](pkg/version.go) and then piped into the Makefile and 
imported into the client to print. This isn't the only way to do it, but seemed reasonable.
At some point I'll want to integrate GitHub tags for between official releases.

## Organization

### `/cmd`

Main applications for this project. Each folder within (e.g., `/cmd/scif`) is a command line application. 
We shouldn't  put a lot of code here. Code that could be imported by other libraries should go into `pkg`, and private into `internal`.
I think what I'll do is have a `main.go` function that imports and invokes the code from `/pkg`. I've noticed that
it's also common to have an "internal" directory inside of cmd, with subfolders corresponding to
the top level folders of command. For example:

```
/cmd/scif/cli.go           (-this is the scif client, but it's a pretty empty file.
/cmd/internal/cli/scif.go  (-most of its functionality is imported from here
```

I'm not sure if this internal folder is respected (.e.g., treated as private) akin to
the top level one.

### `/internal`

This is where I'll put private code (that others can't import).

 - `/internal/app/scif` will be private code for scif
 - `/internal/pkg` will be code/libraries shared by the above (maybe utils?)
 

### `/pkg`

Library code that is okay to be used by external applications. I'd probably want to put some of the scif main functions here.

### `/vendor`

Vendor dependencies, commonly managed with  [`dep`](https://github.com/golang/dep).
(I don't think that I should commit these?)

### `/docs`

User docs (for client help)

### `/assets`

Repository assets (images, logo, etc.)

## Not Needed or Used

I don't think I have:

 - web (web assets)
 - configs (configuration files)
 - init (system init)
 - deployments: docker-compose, kubernetes/helm, etc.
 - tools: supporting tools

## Service Application Directories

### `/api`

There might be a link to the specification here.


### `/examples`

Examples for using SCIF.

### `/scripts`

I might have some scripts to perform various build, install, analysis, etc operations.
The goal with these scripts is to keep the root level Makefile small and simple (e.g., `https://github.com/hashicorp/terraform/blob/master/Makefile`).

### `/build`

Packaging and Continuous Integration - this is where a Docker container, and any packages would go (`/build/package`).
CI (travis, circle) go in a `/build/ci` directory. I guess I'm supposed to link to .circleci from here.

### `/test`

Data and scripts for testing. If I need data, I should put in `test/data` or `test/testdata`.
