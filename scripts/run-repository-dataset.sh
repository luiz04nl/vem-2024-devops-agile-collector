#!/usr/bin/env bash

SOURCE_ENV="${PWD}/.env" &&

. $SOURCE_ENV &&
export $(cut -d= -f1 $SOURCE_ENV) &&

export CGO_ENABLED=1

cd src/repository-dataset
go run . > ../../json.json 2> ../../json.json

# You do not have permission to view repository collaborators

