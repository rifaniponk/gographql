#!/bin/bash

docker-compose run bayarin.go ${*}
docker-compose run bayarin.go chown -R $(id -u):$(id -u) migrations/
