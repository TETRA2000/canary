#!/usr/bin/env bash

set -e

BASE_IMAGE="canary-for-testing"
TEST_IMAGE="canary-test"
TEST_DOCKERFILE="Dockerfile.test"

docker build -t ${BASE_IMAGE} --build-arg USE_HOST_VENDOR=1 .
docker build -f ${TEST_DOCKERFILE} -t ${TEST_IMAGE} .
docker rmi ${TEST_IMAGE} > /dev/null
docker rmi ${BASE_IMAGE} > /dev/null
