#!/usr/bin/env bash

set -e

TEST_IMAGE="canary-test"
TEST_DOCKERFILE="Dockerfile.test"

docker build -t canary . > /dev/null
docker build -f ${TEST_DOCKERFILE} -t ${TEST_IMAGE} .
docker rmi ${TEST_IMAGE} > /dev/null
