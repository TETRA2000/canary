#!/usr/bin/env bash

set -ue

for path in $( ls -1d plugins/*/ ); do
    PLUGIN_NAME=$(basename $path)
    echo Building $(basename $path) plugin...
    go build -buildmode=plugin  -o plugins/${PLUGIN_NAME}.so plugins/${PLUGIN_NAME}/${PLUGIN_NAME}.go
done
