# /pkg

This folder includes scif library code that is okay for use by external applications.
The idea is that other projects will import these libraries and expect them to work.
Please open an issue if something doesn't work.

 - [client](client): is the scif client. You can see how it's interacted with via the [scif entrypoint command](../cmd/scif)
 - [util](util): is various utility functions for scif
 - [version](version): is nothing exciting... the version string!
