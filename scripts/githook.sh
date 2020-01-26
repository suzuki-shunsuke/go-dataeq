#!/usr/bin/env bash

set -eu

ee() {
  echo "+ $*"
  eval "$@"
}

if [ ! -f .git/hooks/pre-commit ]; then
  ee ln -s ../../githooks/pre-commit.sh .git/hooks/pre-commit
fi

ee chmod a+x githooks/*
