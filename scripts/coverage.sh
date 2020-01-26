#!/usr/bin/env bash

set -eu

cd "$(dirname "$0")/.."

go test ./... -race -coverprofile=.coverage.txt -covermode=atomic
go tool cover -html=.coverage.txt
