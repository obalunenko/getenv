#!/bin/bash

set -Eeuo pipefail

SCRIPT_NAME="$(basename "$0")"
SCRIPT_DIR="$(dirname "$0")"
REPO_ROOT="$(cd "${SCRIPT_DIR}" && git rev-parse --show-toplevel)"
SCRIPTS_DIR="${REPO_ROOT}/scripts"
COVER_DIR=${REPO_ROOT}/coverage

source "${SCRIPTS_DIR}/helpers-source.sh"

echo "${SCRIPT_NAME} is running... "

rm -rf "${COVER_DIR}"
mkdir -p "${COVER_DIR}"

go test --count=1 -coverprofile "${COVER_DIR}/full.cov" -covermode=atomic ./...

echo "${SCRIPT_NAME} is done... "
