#!/usr/bin/env bash
set -e

SOURCE_ENV="${PWD}/.env" &&

. $SOURCE_ENV &&
export $(cut -d= -f1 $SOURCE_ENV) &&

export CGO_ENABLED=1

pythonOutputDirectory="out/code-metrics"
if [ -d "$pythonOutputDirectory" ]; then
    rm -r $pythonOutputDirectory
fi
mkdir -p $pythonOutputDirectory

cd src/code-metrics

# OBS https://github.com/mauricioaniche/ck

MSYS_NO_PATHCONV=1 docker build -t ck -f docker/ck.Dockerfile . 2>&1

go run . > ../../out/out-code-metrics.json 2> ../../out/out-code-metrics.json

