@ECHO OFF

SET BASE_IMAGE="canary-for-testing"
SET TEST_IMAGE="canary-test"
SET TEST_DOCKERFILE="Dockerfile.test"

docker build -t %BASE_IMAGE% . || exit /b
docker build -f %TEST_DOCKERFILE% -t %TEST_IMAGE% . || exit /b
docker rmi %TEST_IMAGE% > NUL
docker rmi %BASE_IMAGE% > NUL
