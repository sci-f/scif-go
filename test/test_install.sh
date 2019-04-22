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

# Helper functions
function DirExists {
     if [ -d "${1}" ]; then 
         return 0
     fi
     return 1
}

function FileExists {
     if [ -f "${1}" ]; then 
         return 0
     fi
     return 1
}

# Create temporary scif
SCIF_BASE=$(mktemp -d -t scif.XXXXX)
export SCIF_BASE
sciftest "INSTALL recipe install" 0 ${SCIF} install ../hello-world.scif
sciftest "INSTALL failed install" 255 ${SCIF} install || true
sciftest "INSTALL apps folder exists" 0 DirExists "${SCIF_BASE}/apps/hello-custom"
sciftest "INSTALL metadata folder exists" 0 DirExists "${SCIF_BASE}/apps/hello-custom/scif"
sciftest "INSTALL bin exists" 0 DirExists "${SCIF_BASE}/apps/hello-custom/bin"
sciftest "INSTALL lib exists" 0 DirExists "${SCIF_BASE}/apps/hello-custom/lib"
sciftest "INSTALL runscript exists" 0 FileExists "${SCIF_BASE}/apps/hello-custom/scif/runscript"
sciftest "INSTALL recipe exists" 0 FileExists "${SCIF_BASE}/apps/hello-custom/scif/hello-custom.scif"
sciftest "INSTALL bin script exists" 0 FileExists "${SCIF_BASE}/apps/hello-world-script/bin/hello-world.sh"
sciftest "INSTALL runscript exists" 0 FileExists "${SCIF_BASE}/apps/hello-world-env/scif/labels.json"
sciftest "INSTALL environment exists" 0 FileExists "${SCIF_BASE}/apps/hello-world-env/scif/environment.sh"

rm -rf ${SCIF_BASE}
unset SCIF_BASE
