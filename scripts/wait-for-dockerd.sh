#!/bin/sh
# wait-for-dockerd.sh

set -e

host="$1"
shift
cmd="$@"

until docker info; do
  >&2 echo "Docker daemon is unavailable - sleeping"
  sleep 1
done

>&2 echo "Docker daemon is up - executing command"
exec $cmd
