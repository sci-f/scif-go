# Documentation

Here you can read about usage, development tips, and a quick (docker) tutorial.

 - [usage](usage.md)
 - [development](development.md)
 - [development story](story.md)
 - [organization](organization.md)
 - [tutorial](tutorial.md)


## Frequently Asked Questions

1. Should I use GoLang or Python?

If you want an interactive experience (e.g., a Python shell with a client you
can inspect environments and interact directly with a scientific filesystem and/or
its functions, then you will be happier to use the [scif python](https://www.github.com/vsoch/scif)
base. If you want a ready to go binary and don't care about interaction, or want
to develop with GoLang, you're in the right spot.


2. What containers are available?

If you want to see an installed scif in a Docker Container, see the [tutorial](tutorial.md).
If you want a client base (e.g., the scif executable installed in a container without
development libraries) then pull:

```bash
$ docker pull vanessa/scif-go:latest
```

If you want a container for development, then you should [build it](development.md).
And generally, if you use a provided container, pull it based on the commit
or version (release) tag because latest is a moving target. See [Docker Hub](https://hub.docker.com/r/vanessa/scif-go/tags)
for details.
