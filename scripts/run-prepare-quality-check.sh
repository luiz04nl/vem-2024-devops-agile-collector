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

docker-compose -f sonar-docker-compose.yml down

rm -rf docker/sonarqube
rm -rf docker/postgres

mkdir -p docker/sonarqube
mkdir -p docker/postgres

docker-compose -f sonar-docker-compose.yml up -d

sleep 20

