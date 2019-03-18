#!/usr/bin/env bash

set -e

TEST_IMAGE="canary-test"
TEST_DOCKERFILE="Dockerfile.test"

docker build -f ${TEST_DOCKERFILE} -t ${TEST_IMAGE} --build-arg USE_HOST_VENDOR=1 .
docker rmi ${TEST_IMAGE} > /dev/null
