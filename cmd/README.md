# cmd

This is where you find the main application for the project, the scif binary,
which is in the folder [scif](scif) and tecnically is called the main package.
The functions here invoke others under [pkg/client](../pkg/client) that are
expected to be used by other calling libraries. I didn't follow the standard
practice to "keep this folder really tiny and call things from elsewhere"
because I expected to find the functions here.
