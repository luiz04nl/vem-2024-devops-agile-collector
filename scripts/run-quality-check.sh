#!/usr/bin/env bash
set -e

SOURCE_ENV="${PWD}/.env" &&

. $SOURCE_ENV &&
export $(cut -d= -f1 $SOURCE_ENV) &&

export CGO_ENABLED=1

pythonOutputDirectory="out/quality-check-repos"
if [ -d "$pythonOutputDirectory" ]; then
    rm -r $pythonOutputDirectory
fi
mkdir -p $pythonOutputDirectory

cd src/quality-check-repos

# docker-compose -f sonar-docker-compose.yml down
# rm -r docker
# Set the sonar username and password to admin / sonar
mkdir -p docker/sonarqube
mkdir -p docker/postgres
docker-compose -f sonar-docker-compose.yml up -d
# permission is created automatically: Login: admin. Password: admin.

go run . > ../../out/out-quality-check-repos.json 2> ../../out/out-quality-check-repos.json

