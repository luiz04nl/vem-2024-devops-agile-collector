#!/usr/bin/env bash
set -e

SOURCE_ENV="${PWD}/.env" &&

. $SOURCE_ENV &&
export $(cut -d= -f1 $SOURCE_ENV) &&

export CGO_ENABLED=1

mkdir -p out/create-chart/

cd src/create-charts

# go run . > ../../out/out-create-charts.json 2> ../../out/out-create-charts.json
python main.py > ../../out/out-create-charts.json 2> ../../out/out-create-charts.json
