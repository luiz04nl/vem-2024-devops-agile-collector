#!/usr/bin/env bash
set -e

SOURCE_ENV="${PWD}/.env" &&

. $SOURCE_ENV &&
export $(cut -d= -f1 $SOURCE_ENV) &&

export CGO_ENABLED=1

cd src/check-devops-and-tools

go run . > ../../out/out-check-devops-and-tools.json 2> ../../out/out-check-devops-and-tools.json


