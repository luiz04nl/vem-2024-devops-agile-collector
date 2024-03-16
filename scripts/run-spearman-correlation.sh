#!/usr/bin/env bash
set -e

SOURCE_ENV="${PWD}/.env" &&

. $SOURCE_ENV &&
export $(cut -d= -f1 $SOURCE_ENV) &&

export CGO_ENABLED=1

mkdir -p ./out/spearman-correlation

cd src/spearman-correlation

python main.py > ../../out/out-spearman-correlation.json 2> ../../out/out-spearman-correlation.json


