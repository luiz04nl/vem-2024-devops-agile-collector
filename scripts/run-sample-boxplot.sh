#!/usr/bin/env bash
set -e

SOURCE_ENV="${PWD}/.env" &&

. $SOURCE_ENV &&
export $(cut -d= -f1 $SOURCE_ENV) &&

export CGO_ENABLED=1

mkdir -p out/sample-boxplot

cd src/sample-boxplot

# OBS: inicial / mediana / teceiro quartial / autilier

python main.py > ../../out/out-sample-boxplot.json 2> ../../out/out-sample-boxplot.json


