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
sciftest "INSPECT runscript" 0 ${SCIF} inspect -r hello-custom
sciftest "INSPECT labels" 0 ${SCIF} inspect -l hello-world-env
sciftest "INSPECT install" 0 ${SCIF} inspect -i hello-world-script
sciftest "INSPECT all" 0 ${SCIF} inspect -a hello-world-script
sciftest "INSPECT files" 0 ${SCIF} inspect -f hello-world-script

rm -rf ${SCIF_BASE}
unset SCIF_BASE
