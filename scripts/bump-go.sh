#!/usr/bin/env bash

set -euo pipefail

readonly CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
readonly ROOT_DIR="$(dirname "$CURRENT_DIR")"
# The root go.mod holds the minimum Go version for library users and is bumped
# by hand. This script bumps the development toolchain: tool modules, CI, docs.
readonly GO_MOD_FILE="${ROOT_DIR}/tools/golangci-lint/go.mod"

function main() {
  local input_go_version="${1:-}"
  if [[ -z "${input_go_version}" ]]; then
    echo "Usage: $(basename "$0") <go-version>"
    echo "Examples:"
    echo "  $(basename "$0") 1.25"
    echo "  $(basename "$0") 1.25.x"
    exit 1
  fi

  echo "Updating Go version:"

  local current_go_version
  current_go_version="$(extract_current_version)"
  echo " - Current: ${current_go_version}"
  local escaped_current_go_version
  escaped_current_go_version="$(escape_for_sed_pattern "${current_go_version}")"

  local normalized_go_version
  normalized_go_version="$(normalize_go_version "${input_go_version}")"
  local patched_go_version
  patched_go_version="$(to_patched_go_version "${normalized_go_version}")"
  local ci_go_version
  ci_go_version="${patched_go_version}"
  local toolchain_go_version
  toolchain_go_version="$(to_toolchain_go_version "${patched_go_version}")"

  local escaped_normalized_go_version
  escaped_normalized_go_version="$(escape_for_sed_replacement "${patched_go_version}")"
  local escaped_ci_go_version
  escaped_ci_go_version="$(escape_for_sed_replacement "${ci_go_version}")"
  local escaped_toolchain_go_version
  escaped_toolchain_go_version="$(escape_for_sed_replacement "${toolchain_go_version}")"

  echo " - New (go.mod): ${patched_go_version}"
  echo " - New (toolchain): ${toolchain_go_version}"
  echo " - New (CI): ${ci_go_version}"

  # bump mod files in all the modules
  while IFS= read -r mod_file; do
    bump_mod_file "${mod_file}" "${escaped_current_go_version}" "${escaped_normalized_go_version}" "${escaped_toolchain_go_version}"
  done < <(find "${ROOT_DIR}" -name "go.mod" -not -path "${ROOT_DIR}/go.mod" -not -path "${ROOT_DIR}/vendor/*" -not -path "${ROOT_DIR}/.git/*")

  # bump markdown files
  while IFS= read -r f; do
    bump_golang_docker_images "${f}" "${escaped_current_go_version}" "${escaped_normalized_go_version}"
  done < <(find "${ROOT_DIR}" -name "*.md")

  # bump github action workflows
  while IFS= read -r f; do
    bump_ci_matrix "${f}" "${escaped_ci_go_version}"
  done < <(find "${ROOT_DIR}/.github/workflows" -name "*.yml")
}

function escape_for_sed_pattern() {
  echo "${1}" | sed -e 's/[.[\*^$()+?{}|]/\\&/g'
}

function escape_for_sed_replacement() {
  echo "${1}" | sed -e 's/[&/]/\\&/g'
}

function normalize_go_version() {
  local version="${1}"
  version="${version#go}"
  version="${version%.x}"

  if ! echo "${version}" | grep -Eq '^[0-9]+\.[0-9]+(\.[0-9]+)?$'; then
    echo "Invalid Go version: ${1}" >&2
    exit 1
  fi

  echo "${version}"
}

function to_patched_go_version() {
  local version="${1}"
  if echo "${version}" | grep -Eq '^[0-9]+\.[0-9]+\.[0-9]+$'; then
    echo "${version}"
    return
  fi

  echo "${version}.0"
}

function to_toolchain_go_version() {
  local version="${1}"
  echo "go${version}"
}

# it will replace matrix entries like 'go-version: [1.24.x]' with the new CI version.
function bump_ci_matrix() {
  local file="${1}"
  local new_go_version="${2}"

  sed -E "s/(go-version:[[:space:]]*\[)[^]]+(\])/\1${new_go_version}\2/g" "${file}" > "${file}.tmp"
  mv "${file}.tmp" "${file}"
}

# it will replace the 'golang:${old_go_version}' with 'golang:${new_go_version}' in the given file
function bump_golang_docker_images() {
  local file="${1}"
  local old_go_version="${2}"
  local new_go_version="${3}"

  sed "s/golang:${old_go_version}/golang:${new_go_version}/g" "${file}" > "${file}.tmp"
  mv "${file}.tmp" "${file}"

}

# it will replace 'go ${old_go_version}' and 'toolchain go...' with the new values in go.mod
function bump_mod_file() {
  local go_mod_file="${1}"
  local old_go_version="${2}"
  local new_go_version="${3}"
  local new_toolchain_go_version="${4}"

  sed -E \
    -e "s/^go ${old_go_version}$/go ${new_go_version}/g" \
    -e "s/^toolchain go[0-9]+\.[0-9]+(\.[0-9]+)?$/toolchain ${new_toolchain_go_version}/g" \
    "${go_mod_file}" > "${go_mod_file}.tmp"
  mv "${go_mod_file}.tmp" "${go_mod_file}"

}

# This function reads the tool go.mod file and extracts the current Go version.
function extract_current_version() {
  cat "${GO_MOD_FILE}" | grep '^go .*' | sed 's/^go //g' | head -n 1
}

main "$@"
