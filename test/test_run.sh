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

# Create temporary scif
SCIF_BASE=$(mktemp -d -t scif.XXXXX)
export SCIF_BASE
sciftest "INSTALL recipe install" 0 ${SCIF} install ../hello-world.scif
sciftest "RUN hello-custom" 0 ${SCIF} run hello-custom
sciftest "RUN hello-world-script" 0 ${SCIF} run hello-world-script

rm -rf ${SCIF_BASE}
unset SCIF_BASE
