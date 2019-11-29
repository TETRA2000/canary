#!/usr/bin/env bash
set -e
gofmt -s -l $(find . -type f -name '*.go'| grep -v "/vendor/")
