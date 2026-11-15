#!/bin/bash

set -Eeuo pipefail

function checkInstalled() {
  CMD=$1

  if ! command -v "${CMD}" &>/dev/null; then
    echo "Install ${CMD}"
    echo "run 'make install-tools' "
    exit 1
  fi
}

# Prints the binary path of a tool pinned in tools/<name>/go.mod, building it on
# first use. go tool runs inside the tool module: -modfile from the repo root
# clashes with vendor/.
function toolPath() {
  local name="$1"

  go tool -C "${REPO_ROOT}/tools/${name}" -n "${name}"
}

# Runs a tool pinned in tools/<name>/go.mod: gotool <name> [args...]
function gotool() {
  local name="$1"
  shift

  local bin
  bin="$(toolPath "${name}")"
  "${bin}" "$@"
}

function openSource() {
  OPEN_CMD=""
  OPEN_SRC=$1

  case "$OSTYPE" in
  darwin*)
    OPEN_CMD=open
    ;;
  linux*)
    OPEN_CMD=xdg-open
    ;;
  msys* | cygwin*)
    OPEN_CMD=start
    ;;
  *)
    echo "unknown: $OSTYPE"
    exit 1
    ;;
  esac

  ${OPEN_CMD} "${OPEN_SRC}"
}
