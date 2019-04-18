# Development Notes

These are notes that I took (from an original template README.md) that I (somewhat)
rewrote in my own words, mainly to understand what was going on. Once I had developed
a bit and realized that the examples provided were excessive, I pruned a lot of
this away. See [development](development.md) documentation for the final notes.

## Things I Changed My Mind About

### Floating Docs Package

> there should be programmatic docs in a (top level) docs folder that is imported in cmd/internal/cli/scif.go (the help strings, etc.)

This was a pattern I saw a few times, and I realized that (for the most part) this "docs" package was
providing templates and content for the scif command (under scif/cmd). Thus, it was more logical to me
to put it into scif/cmd/docs and have it be a part of the main package.

### Minimal cmd Folder

Most libraries put main applications as subfolders under "cmd." They also mentioned
that 

> you shouldn't put a lot of code here. Code that could be imported by other libraries 
should go into `pkg`, and private into `internal`.

It was logical that shared libraries should be provided as a package. But I saw a lot
of weird patterns of having the command client code stuck in weird places like:

```bash
/cmd/scif/cli.go           (-this is the scif client, but it's a pretty empty file.
/cmd/internal/cli/scif.go  (-most of its functionality is imported from here
```

And for the life of me, I could be working on something for a while and still look
in the wrong spot for it. I fundamentally disagree with the idea of putting next
to nothing in `cmd/<app>` folders, because this is *exactly* where you would
expect to find the client (mostly cobra) based functions for flags, etc.
I looked at a lot of the [containers](https://www.github.com/containers)
organization repos on GitHub, and they tend to follow what I think is more intuitive.
Take a look at [Podman](https://github.com/containers/libpod/tree/master/cmd/podman).
You can find all the entrypoint commands easily (by name).


## Things I Agreed With

### The Logger Package can be Internal

> internal/pkg/logger should have the logger

I found it generally strange that such effort was placed on making many of
the functions internal. It would even lead to weird structure like:

```
cmd/scif
cmd/internal/scif
```

What could be so necessary and secret to have these confusing two locations for the
same thing? I didn't want to do this. However, for the logger, this was a clear
package that had no good reason to be exposed.

### Version Variable

Most packages would have some version.go (or similar) and then pipe it where it needed
to go. This is reasonable to me, along with maybe having a GitHub commit somewhere 
for the more specific versioning.

## General Notes

This was rewritten in my own words, mostly for learning. I didn't follow or use
most of it.

### `/internal`

This is where I'll put private code (that others can't import).

 - `/internal/app/scif` will be private code for scif
 - `/internal/pkg` will be code/libraries shared by the above (maybe utils?)
 
### `/pkg`

Library code that is okay to be used by external applications. 
I'd probably want to put some of the scif main functions here that get called
by the command (note that I did do this!)

### `/vendor`

Vendor dependencies, commonly managed with  [`dep`](https://github.com/golang/dep).
(I don't think that I should commit these?)

### `/docs`

User docs (for client help)

### `/assets`

Repository assets (images, logo, etc.)

### Not Needed or Used

I don't think I have:

 - web (web assets)
 - configs (configuration files)
 - init (system init)
 - deployments: docker-compose, kubernetes/helm, etc.
 - tools: supporting tools

### `/api`

There might be a link to the specification here?

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
