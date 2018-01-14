#!/usr/bin/env bash

set -e

if [ "$USE_HOST_VENDOR" == "1" ]; then
    exit
fi

if [ -d ./vendor ] && [ "$(ls -A ./vendor)" ] ; then
    echo "Clearing vendor directory"
    rm -rf ./vendor
fi

echo "Running dep ensure command."
dep ensure
