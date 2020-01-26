#!/usr/bin/env bash

set -eu
set -o pipefail

cd "$(dirname "$0")/.."

find . \
  -type d -name .git -prune -o \
  -type f -name "*.go" -print0 |
  xargs -0 gofmt -l -s -w
