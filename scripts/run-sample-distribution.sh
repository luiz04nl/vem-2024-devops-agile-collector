#!/usr/bin/env bash
set -e

SOURCE_ENV="${PWD}/.env" &&

. $SOURCE_ENV &&
export $(cut -d= -f1 $SOURCE_ENV) &&

export CGO_ENABLED=1

mkdir -p out/sample-distribution

cd src/sample-distribution

python main.py > ../../out/out-sample-distribution.json 2> ../../out/out-sample-distribution.json

# Outras ferramentas ja informacao essa metrica

# OBS: tentar contar linhas de codigo com cloc
# https://github.com/AlDanial/cloc
# https://cloc.sourceforge.net/#:~:text=cloc%20counts%20blank%20lines%2C%20comment,%2C%20comment%2C%20and%20source%20lines.
