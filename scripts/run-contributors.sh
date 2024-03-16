#!/usr/bin/env bash
set -e

SOURCE_ENV="${PWD}/.env" &&

. $SOURCE_ENV &&
export $(cut -d= -f1 $SOURCE_ENV) &&

export CGO_ENABLED=1

pythonOutputDirectory="out/contributors"
if [ -d "$pythonOutputDirectory" ]; then
    rm -r $pythonOutputDirectory
fi
mkdir -p $pythonOutputDirectory

cd src/contributors

go run . > ../../out/out-contributors.json 2> ../../out/out-contributors.json
