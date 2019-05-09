#!/usr/bin/env bash

command="$1"
version="v1.16.0"

if [[ "${command}" = "install" ]]; then
    go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
elif [[ "${command}" = "run" ]]; then
    golangci-lint run
fi
