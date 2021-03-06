#!/bin/sh

set -eu pipefail

SCRIPT_NAME="$(basename "$0")"

echo "${SCRIPT_NAME} is running... "

export AOC_REGRESSION_ENABLED=true

GOTEST="go test -v -race"
if [ -f "$(go env GOPATH)/bin/gotestsum" ] || [ -f "/usr/local/bin/gotestsum" ]; then
  GOTEST="gotestsum --format pkgname --"
fi

unset AOC_REGRESSION_ENABLED

${GOTEST} ./...

echo "${SCRIPT_NAME} done."
