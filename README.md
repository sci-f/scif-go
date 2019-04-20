# SCI-F GoLang

This is the GoLang implementation of the Scientific Filesystem, and the first full GoLang package
that I'm developing. 

[![CircleCI](https://circleci.com/gh/sci-f/scif-go.svg?style=svg)](https://circleci.com/gh/sci-f/scif-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/sci-f/scif-go?style=flat-square)](https://goreportcard.com/report/github.com/sci-f/scif-go)

## Development

### Commands

If you want to install dependencies (golint):

```bash
$ make deps
```

To build the package (the scif client goes into the [bin](bin) folder)

```bash
$ make build
```

To format all the files pretty:

```bash
$ make fmt
```

### Documentation

 - [Documentation for Scif](docs) including development, history, and organization
 - [Docker](docker) including instructions for building development (and tiny) Docker containers for scif
