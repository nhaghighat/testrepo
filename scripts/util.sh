#!/bin/bash

# This script installs a multi-node kubernetes cluster with minikube for testing purposes

set -euo pipefail

log() {
  timestamp=$(date +"[%m%d %H:%M:%S]")
  echo "${timestamp} ${1-}" >&2
  shift
  for message; do
    echo "    ${message}" >&2
  done
}

# looks for $1 in the $PATH
find-binary() {
  if which "$1" >/dev/null; then
    echo -n "${1}"
  fi
}