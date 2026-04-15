#!/usr/bin/env bash

set -euo pipefail

readonly CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
readonly ROOT_DIR="$(dirname "$CURRENT_DIR")"
readonly GO_MOD_FILE="${ROOT_DIR}/go.mod"

function main() {
  local inputGoVersion="${1:-}"
  if [[ -z "${inputGoVersion}" ]]; then
    echo "Usage: $(basename "$0") <go-version>"
    echo "Examples:"
    echo "  $(basename "$0") 1.25"
    echo "  $(basename "$0") 1.25.x"
    exit 1
  fi

  echo "Updating Go version:"

  local currentGoVersion
  currentGoVersion="$(extractCurrentVersion)"
  echo " - Current: ${currentGoVersion}"
  local escapedCurrentGoVersion
  escapedCurrentGoVersion="$(escapeForSedPattern "${currentGoVersion}")"

  local normalizedGoVersion
  normalizedGoVersion="$(normalizeGoVersion "${inputGoVersion}")"
  local patchedGoVersion
  patchedGoVersion="$(toPatchedGoVersion "${normalizedGoVersion}")"
  local ciGoVersion
  ciGoVersion="${patchedGoVersion}"
  local toolchainGoVersion
  toolchainGoVersion="$(toToolchainGoVersion "${patchedGoVersion}")"

  local escapedNormalizedGoVersion
  escapedNormalizedGoVersion="$(escapeForSedReplacement "${patchedGoVersion}")"
  local escapedCIGoVersion
  escapedCIGoVersion="$(escapeForSedReplacement "${ciGoVersion}")"
  local escapedToolchainGoVersion
  escapedToolchainGoVersion="$(escapeForSedReplacement "${toolchainGoVersion}")"

  echo " - New (go.mod): ${patchedGoVersion}"
  echo " - New (toolchain): ${toolchainGoVersion}"
  echo " - New (CI): ${ciGoVersion}"

  # bump mod files in all the modules
  while IFS= read -r modFile; do
    bumpModFile "${modFile}" "${escapedCurrentGoVersion}" "${escapedNormalizedGoVersion}" "${escapedToolchainGoVersion}"
  done < <(find "${ROOT_DIR}" -name "go.mod" -not -path "${ROOT_DIR}/vendor/*" -not -path "${ROOT_DIR}/.git/*")

  # bump markdown files
  while IFS= read -r f; do
    bumpGolangDockerImages "${f}" "${escapedCurrentGoVersion}" "${escapedNormalizedGoVersion}"
  done < <(find "${ROOT_DIR}" -name "*.md")

  # bump github action workflows
  while IFS= read -r f; do
    bumpCIMatrix "${f}" "${escapedCIGoVersion}"
  done < <(find "${ROOT_DIR}/.github/workflows" -name "*.yml")
}

function escapeForSedPattern() {
  echo "${1}" | sed -e 's/[.[\*^$()+?{}|]/\\&/g'
}

function escapeForSedReplacement() {
  echo "${1}" | sed -e 's/[&/]/\\&/g'
}

function normalizeGoVersion() {
  local version="${1}"
  version="${version#go}"
  version="${version%.x}"

  if ! echo "${version}" | grep -Eq '^[0-9]+\.[0-9]+(\.[0-9]+)?$'; then
    echo "Invalid Go version: ${1}" >&2
    exit 1
  fi

  echo "${version}"
}

function toPatchedGoVersion() {
  local version="${1}"
  if echo "${version}" | grep -Eq '^[0-9]+\.[0-9]+\.[0-9]+$'; then
    echo "${version}"
    return
  fi

  echo "${version}.0"
}

function toToolchainGoVersion() {
  local version="${1}"
  echo "go${version}"
}

# it will replace matrix entries like 'go-version: [1.24.x]' with the new CI version.
function bumpCIMatrix() {
  local file="${1}"
  local newGoVersion="${2}"

  sed -E "s/(go-version:[[:space:]]*\[)[^]]+(\])/\1${newGoVersion}\2/g" "${file}" > "${file}.tmp"
  mv "${file}.tmp" "${file}"
}

# it will replace the 'golang:${oldGoVersion}' with 'golang:${newGoVersion}' in the given file
function bumpGolangDockerImages() {
  local file="${1}"
  local oldGoVersion="${2}"
  local newGoVersion="${3}"

  sed "s/golang:${oldGoVersion}/golang:${newGoVersion}/g" "${file}" > "${file}.tmp"
  mv "${file}.tmp" "${file}"

}

# it will replace 'go ${oldGoVersion}' and 'toolchain go...' with the new values in go.mod
function bumpModFile() {
  local goModFile="${1}"
  local oldGoVersion="${2}"
  local newGoVersion="${3}"
  local newToolchainGoVersion="${4}"

  sed -E \
    -e "s/^go ${oldGoVersion}$/go ${newGoVersion}/g" \
    -e "s/^toolchain go[0-9]+\.[0-9]+(\.[0-9]+)?$/toolchain ${newToolchainGoVersion}/g" \
    "${goModFile}" > "${goModFile}.tmp"
  mv "${goModFile}.tmp" "${goModFile}"

}

# This function reads the root go.mod file and extracts the current Go version.
function extractCurrentVersion() {
  cat "${GO_MOD_FILE}" | grep '^go .*' | sed 's/^go //g' | head -n 1
}

main "$@"
