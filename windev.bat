@echo off
docker-compose build && docker-compose down & docker volume rm canary_dockerd-run-volume & docker-compose up
