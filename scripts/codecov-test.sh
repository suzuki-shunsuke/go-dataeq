#!/usr/bin/env bash
# https://github.com/codecov/example-go#caveat-multiple-files

set -eu

cd "$(dirname "$0")/.."

go test -race -coverprofile=coverage.out -covermode=atomic ./dataeq

test "${LOCAL:-false}" = "true" || bash <(curl -s https://codecov.io/bash)
