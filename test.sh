#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

echo "run tests:"
go test ./pkg/...
echo

echo -n "check gofmt: "
ERRORS=$(find pkg -type f -name \*.go | xargs gofmt -l 2>&1 || true)
if [ -n "${ERRORS}" ]; then
    echo "these files need to be gofmt'ed:"
    for error in ${ERRORS}; do
        echo "    $error"
    done
    echo "FAIL"
    exit 1
fi
echo "PASS"
echo

echo -n "check go vet: "
ERRORS=$(go vet ./pkg/... 2>&1 || true)
if [ -n "${ERRORS}" ]; then
    echo "${ERRORS}"
    echo "FAIL"
    exit 1
fi
echo "PASS"
echo