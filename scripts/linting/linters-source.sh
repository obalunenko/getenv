#!/bin/bash

set -Eeuo pipefail

SCRIPT_DIR="$(dirname "$0")"
REPO_ROOT="$(cd "${SCRIPT_DIR}" && git rev-parse --show-toplevel)"
SCRIPTS_DIR="${REPO_ROOT}/scripts"

source "${SCRIPTS_DIR}/helpers-source.sh"

function vet() {
  echo "vet project..."

  # Rely on the exit code: stderr also carries non-findings such as "go: downloading ...".
  if ! go vet ./...; then
    echo "fix it"
    exit 1
  fi

  echo "code is ok"
  echo ""
}

function fmt() {
  echo "fmt lint..."

  check_installed 'gofmt'

  declare -a fmts=$(gofmt -s -l $(find . -type f -name '*.go' | grep -v 'vendor' | grep -v '.git'))

  if [[ ${fmts} ]]; then
    echo "fix it:"
    for f in "${fmts[@]}"; do
      echo "$f"

    done
    exit 1

  else
    echo "code is ok"
    echo ${fmts}
  fi
  echo ""
}

function go_group() {
  echo "goimports..."

  declare -a lints=$(go_tool goimports -l -local=$(go list -m) $(find . -type f -name "*.go" | grep -v "vendor/"))

  if [[ ${lints} ]]; then
    echo "fix it:"
    for l in "${lints[@]}"; do
      echo "$l"

    done
    exit 1

  else
    echo "code is ok"
    echo ${lints}
  fi
  echo ""

}

function golangci() {
  echo "golangci-lint linter running..."

  go_tool golangci-lint run --config .golangci.yml ./...

  echo ""
}
