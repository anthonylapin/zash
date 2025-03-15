#!/bin/sh

# Use this script to run program locally.

set -e # Exit early if any commands fail

(
  cd "$(dirname "$0")" # Ensure compile steps are run within the repository directory
  go build -o /tmp/shell-target src/*.go
)

exec /tmp/shell-target

