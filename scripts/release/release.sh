#!/bin/bash

set -Eeuo pipefail

SCRIPT_NAME="$(basename "$0")"
SCRIPT_DIR="$(dirname "$0")"
REPO_ROOT="$(cd "${SCRIPT_DIR}" && git rev-parse --show-toplevel)"
SCRIPTS_DIR="${REPO_ROOT}/scripts"

source "${SCRIPTS_DIR}/helpers-source.sh"

checkInstalled goreleaser

echo "${SCRIPT_NAME} is running... "

# Get new tags from the remote
git fetch --tags -f

export BUILD_COMMIT=$(git rev-parse HEAD)
export BUILD_SHORTCOMMIT=$(git rev-parse --short HEAD)
export BUILD_DATE=$(date -u +%Y-%m-%dT%H:%M:%SZ)
export BUILD_VERSION=$(git describe --tags  --always $(git rev-list --tags --max-count=1))

export BUILD_APPNAME=bitech-go-shared

if [ -z "${BUILD_VERSION}" ] || [ "${BUILD_VERSION}" = "${BUILD_SHORTCOMMIT}" ]
 then
  BUILD_VERSION="v0.0.0"
fi

export BUILD_GOVERSION=$(go version | awk '{print $3;}')

goreleaser release --rm-dist
