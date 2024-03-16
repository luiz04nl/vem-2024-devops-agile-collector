#!/usr/bin/env bash
set -e

SOURCE_ENV="${PWD}/.env" &&

. $SOURCE_ENV &&
export $(cut -d= -f1 $SOURCE_ENV) &&

export CGO_ENABLED=1

cd src/create-dataset

go run . > ../../out/out-create-dataset.json 2> ../../out/out-create-dataset.json
