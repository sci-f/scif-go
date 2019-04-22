# Tests

Since scif is a command line client, we run some of the tests external to it.
You should first build scif before running these tests.

```bash
$ make build
```

And then these tests are run with the entire test suite:

```bash
make test
```

If you want to run all the bash tests (without the go tests) from the root
of the repo (after build) do:

```bash
test/run_tests.sh bin/scif
```

You can also run a single test file by changing directory to tests first,
and calling it directly providing the scif binary you want to test.

```bash
$ cd test
$ ./test_help.sh ../bin/scif
```

See the [Makefile](../Makefile) for how the tests are run with `make test`.
