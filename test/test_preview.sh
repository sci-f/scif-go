#!/usr/bin/env bash

# This is the init script for running SCIF tests. Here we create a dummy
# base folder for the tests to interact with.

set -o errexit
set -o nounset
set -o pipefail

# The user must provide the scif binary as first argument
if [ $# -eq 0 ]
  then
    echo "Please provide the scif binary to test as first argument."
    exit 1
fi

SCIF="${1}"

source ./functions

# Ensure that it exists
if [ ! -f "${SCIF}" ]; then
    echo "scif binary ${SCIF} does not exist."
fi

# Test preview
sciftest "PREVIEW Client preview" 255 ${SCIF} preview || true
sciftest "PREVIEW preview recipe" 0 ${SCIF} preview ../hello-world.scif
