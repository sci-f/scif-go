#!/usr/bin/env bash

# This is the init script for running SCIF tests. Here we create a dummy
# base folder for the tests to interact with.

set -o errexit
set -o nounset
set -o pipefail

echo
echo "Starting Bash Client Testing ----------------"

# Make sure to change directory and back
CURRENT="${PWD}"
HERE="$( cd "$(dirname "$0")" ; pwd -P )"
cd "${HERE}"

# Ensure that the testing scif (bin/scif) exists
SCIF=../bin/scif
if [ ! -f "${SCIF}" ]; then
    echo "scif binary ${SCIF} does not exist."
fi

# Run all tests
for test_file in $( find "${HERE}" -type f -name '*.sh' -not -path '*run_tests.sh' ); do
    echo
    echo "Running test file $(basename ${test_file})"
    /bin/bash "${test_file}" "${SCIF}"
done

# Return to previous directory before run
cd "${CURRENT}"
