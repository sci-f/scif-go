# Docker

You can build Docker containers for development, or for usage of Scif. Note that
the working directory should be the root of the repository (one folder up)
and then the files referenced in this folder:

## Multistage Build

The multistage build is a tiny image that just builds and keeps the scif executable
(and throws away dependencies). You can't use this for development. From the root
of the repository, run:

```bash
$ docker build -f docker/Dockerfile -t vanessa/scif-go .
```

## Development Build

This version includes all dependencies, and is good for development. From the
root of the repository, do:

```bash
$ docker build -f docker/Dockerfile.dev -t vanessa/scif-go .
```
