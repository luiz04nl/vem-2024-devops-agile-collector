#!/usr/bin/env bash
set -e

SOURCE_ENV="${PWD}/.env" &&

. $SOURCE_ENV &&
export $(cut -d= -f1 $SOURCE_ENV) &&

export CGO_ENABLED=1

pythonOutputDirectory="out/code-smell-frequency"
if [ -d "$pythonOutputDirectory" ]; then
    rm -r $pythonOutputDirectory
fi
mkdir -p $pythonOutputDirectory

cd src/code-smell-frequency

# Code smeells mais frequentes - nomes e quantas vezes aparecem
# -- Dado nossa amostra levantar os 5 principais
# para cada um deles a quantidade -

go run . > ../../out/out-code-smell-frequency.json 2> ../../out/out-code-smell-frequency.json
