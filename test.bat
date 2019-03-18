@ECHO OFF

SET TEST_IMAGE="canary-test"
SET TEST_DOCKERFILE="Dockerfile.test"

docker build -f %TEST_DOCKERFILE% -t %TEST_IMAGE% --build-arg USE_HOST_VENDOR=1 . || exit /b
docker rmi %TEST_IMAGE% > NUL
